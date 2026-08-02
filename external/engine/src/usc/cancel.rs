use std::cell::Cell;
use std::sync::atomic::{
    AtomicU8,
    Ordering,
};

thread_local! {
    static CURRENT: Cell<*const AtomicU8> = const { Cell::new(std::ptr::null()) };
}

pub(crate) fn with_token<T>(token: *const AtomicU8, operation: impl FnOnce() -> T) -> T {
    CURRENT.with(|current| {
        let previous = current.replace(token);
        let result = operation();
        current.set(previous);
        result
    })
}

pub(crate) fn check() -> Result<(), String> {
    CURRENT.with(|current| {
        let token = current.get();
        if token.is_null() || unsafe { &*token }.load(Ordering::Relaxed) == 0 {
            Ok(())
        } else {
            Err("engine operation cancelled".to_owned())
        }
    })
}

pub(crate) fn clear() {
    CURRENT.with(|current| current.set(std::ptr::null()));
}
