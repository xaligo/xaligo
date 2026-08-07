use pptx::core_properties::CoreProperties;
use pptx::Presentation;

use crate::cnf::exporter::EMU_PER_INCH;
use crate::ent::model::pptx::{
    DocumentPlan,
    ExporterRequest,
    Op,
    Slide,
};
use crate::rep;
use crate::util::error::Error;

pub fn export(request: &ExporterRequest) -> Result<Vec<u8>, Error> {
    let (pages, legacy, legend, connector_legend) = request.plan.parts();
    let first_slide = pages.first().map(|page| &page.slide).or_else(|| legacy.map(|item| item.0))
        .ok_or_else(|| Error::invalid("PPTX plan must contain at least one page"))?;
    validate_slide(first_slide)?;
    if let DocumentPlan::Current(plan) = &request.plan {
        if plan.schema_version != 2 {
            return Err(Error::invalid("PPTX document plan schemaVersion must be 2"));
        }
        for page in &plan.pages {
            validate_slide(&page.slide)?;
            if page.slide.w != first_slide.w || page.slide.h != first_slide.h {
                return Err(Error::invalid(format!(
                    "PPTX pages must use one slide size; page {} is {}x{}, expected {}x{}",
                    page.id, page.slide.w, page.slide.h, first_slide.w, first_slide.h
                )));
            }
        }
    }

    let mut presentation = Presentation::new()?;
    presentation.set_slide_width(to_emu(first_slide.w))?;
    presentation.set_slide_height(to_emu(first_slide.h))?;
    set_properties(&mut presentation, request)?;
    let layouts = presentation.slide_layouts()?;
    let layout = layouts.first().ok_or_else(|| Error::invalid("PPTX template has no slide layout"))?.clone();

    let mut package_ops = Vec::new();
    for page in pages {
        add_page(&mut presentation, &layout, &page.slide, &page.ops)?;
        package_ops.push(page.ops.clone());
    }
    if let Some((slide, ops)) = legacy {
        add_page(&mut presentation, &layout, slide, ops)?;
        package_ops.push(ops.to_vec());
    }
    package_ops.extend(rep::add_legends(
        &mut presentation,
        &layout,
        first_slide,
        legend,
        connector_legend,
    )?);
    let bytes = presentation.to_bytes()?;
    rep::finalize_package(bytes, &package_ops, request.options.compression.unwrap_or(true))
}

fn add_page(
    presentation: &mut Presentation,
    layout: &pptx::slide::SlideLayoutRef,
    slide: &Slide,
    ops: &[Op],
) -> Result<(), Error> {
    let slide_ref = presentation.add_slide(layout)?;
    let xml = rep::slide_xml(slide, ops)?;
    *presentation.slide_xml_mut(&slide_ref)? = xml.into_bytes();
    Ok(())
}

fn set_properties(presentation: &mut Presentation, request: &ExporterRequest) -> Result<(), Error> {
    let mut properties = CoreProperties::new();
    properties.set_title(request.options.title.as_deref().unwrap_or("xaligo export"));
    properties.set_author(request.options.author.as_deref().unwrap_or("xaligo"));
    properties.set_subject(request.options.subject.as_deref().unwrap_or("xaligo PPTX export"));
    presentation.set_core_properties(&properties)?;
    Ok(())
}

fn validate_slide(slide: &Slide) -> Result<(), Error> {
    if !slide.w.is_finite() || !slide.h.is_finite() || slide.w <= 0.0 || slide.h <= 0.0 {
        return Err(Error::invalid("PPTX page slide size must be positive and finite"));
    }
    Ok(())
}

fn to_emu(inches: f64) -> i64 {
    (inches * EMU_PER_INCH).round() as i64
}
