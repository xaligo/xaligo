---
applyTo: ".github/instructions/manual/**"
---

# 10.13 PPTX and routing: Legend Pages

## Legend Pages

PPTX export adds legend slides after all frame/diagram slides when `--services`
is provided.

- Legend data is derived from `services.csv`.
- Only services actually used in the scene are included.
- The legend contains icon, abbreviation, and official name.
- Legend layout is fixed to 4 columns per slide.
- Additional legend slides may be created when entries exceed one slide.
- The diagram slide should not include an outside-frame legend; the PPTX legend
  belongs on separate slides.
