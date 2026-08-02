pub(crate) struct Decoder<'a> {
    input: &'a [u8],
    offset: usize,
}

impl<'a> Decoder<'a> {
    pub(crate) fn new(input: &'a [u8]) -> Self {
        Self { input, offset: 0 }
    }

    pub(crate) fn read_exact(&mut self, length: usize) -> Result<&'a [u8], String> {
        let end = self
            .offset
            .checked_add(length)
            .ok_or_else(|| "engine request length overflow".to_owned())?;
        if end > self.input.len() {
            return Err("truncated engine request".to_owned());
        }
        let value = &self.input[self.offset..end];
        self.offset = end;
        Ok(value)
    }

    pub(crate) fn read_u8(&mut self) -> Result<u8, String> {
        Ok(self.read_exact(1)?[0])
    }

    pub(crate) fn read_u16(&mut self) -> Result<u16, String> {
        let bytes: [u8; 2] = self
            .read_exact(2)?
            .try_into()
            .map_err(|_| "invalid u16".to_owned())?;
        Ok(u16::from_le_bytes(bytes))
    }

    pub(crate) fn read_u32(&mut self) -> Result<u32, String> {
        let bytes: [u8; 4] = self
            .read_exact(4)?
            .try_into()
            .map_err(|_| "invalid u32".to_owned())?;
        Ok(u32::from_le_bytes(bytes))
    }

    pub(crate) fn read_i32(&mut self) -> Result<i32, String> {
        let bytes: [u8; 4] = self
            .read_exact(4)?
            .try_into()
            .map_err(|_| "invalid i32".to_owned())?;
        Ok(i32::from_le_bytes(bytes))
    }

    pub(crate) fn read_u64(&mut self) -> Result<u64, String> {
        let bytes: [u8; 8] = self
            .read_exact(8)?
            .try_into()
            .map_err(|_| "invalid u64".to_owned())?;
        Ok(u64::from_le_bytes(bytes))
    }

    pub(crate) fn read_f64(&mut self) -> Result<f64, String> {
        let bytes: [u8; 8] = self
            .read_exact(8)?
            .try_into()
            .map_err(|_| "invalid f64".to_owned())?;
        Ok(f64::from_le_bytes(bytes))
    }

    pub(crate) fn is_empty(&self) -> bool {
        self.offset == self.input.len()
    }
}
