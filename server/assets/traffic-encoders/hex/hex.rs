extern crate alloc;
extern crate core;
extern crate hex;

use std::alloc::{alloc, dealloc, Layout};
use std::slice;

fn encode(input: &[u8]) -> Vec<u8> {
    let mut output: Vec<u8> = vec![0; input.len() * 2];
    hex::encode_to_slice(input, &mut output).unwrap();
    return output;
}

fn decode(input: &[u8]) -> Vec<u8> {
    let mut output: Vec<u8> = vec![0; input.len() / 2];
    hex::decode_to_slice(input, &mut output).unwrap();
    return output;
}

#[link(wasm_import_module = "hex")]
extern "C" {
    #[link_name = "log"]
    fn _log(ptr: u32, size: u32);
}

/// WebAssembly export that accepts a string (linear memory offset, byteCount)
/// and returns a pointer/size pair packed into a u64.
///
#[cfg_attr(all(target_arch = "wasm32"), export_name = "encode")]
#[no_mangle]
pub unsafe extern "C" fn _encode(ptr: u32, len: u32) -> u64 {
    let input = slice::from_raw_parts_mut(ptr as *mut u8, len as usize);
    let output = encode(input);
    let (ptr, len) = (output.as_ptr(), output.len());
    // Note: This changes ownership of the pointer to the external caller.
    std::mem::forget(output);
    return ((ptr as u64) << 32) | len as u64;
}

/// WebAssembly export that accepts a string (linear memory offset, byteCount)
/// and returns a pointer/size pair packed into a u64.
///
#[cfg_attr(all(target_arch = "wasm32"), export_name = "decode")]
#[no_mangle]
pub unsafe extern "C" fn _decode(ptr: u32, len: u32) -> u64 {
    let input = slice::from_raw_parts_mut(ptr as *mut u8, len as usize);
    let output = decode(input);
    let (ptr, len) = (output.as_ptr(), output.len());
    // Note: This changes ownership of the pointer to the external caller.
    std::mem::forget(output);
    return ((ptr as u64) << 32) | len as u64;
}

/// Logs a message to the console using [`_log`].
fn log(message: &String) {
    unsafe {
        let (ptr, len) = string_to_ptr(message);
        _log(ptr, len);
    }
}

/// Returns a pointer and size pair for the given string in a way compatible
/// with WebAssembly numeric types.
unsafe fn string_to_ptr(s: &String) -> (u32, u32) {
    return (s.as_ptr() as u32, s.len() as u32);
}

/// WebAssembly export that allocates a pointer (linear memory offset) that can
/// be used for a string.
#[cfg_attr(all(target_arch = "wasm32"), export_name = "malloc")]
#[no_mangle]
pub extern "C" fn _allocate(size: u32) -> *mut u8 {
    unsafe {
        let ptr = allocate(size as usize);
        return ptr;
    }
}

/// Allocates size bytes and leaks the pointer where they start.
unsafe fn allocate(size: usize) -> *mut u8 {
    let layout = Layout::from_size_align(size, std::mem::align_of::<u8>()).expect("Bad layout");
    let ptr = alloc(layout);
    return ptr;
}

/// WebAssembly export that deallocates a pointer of the given size.
#[cfg_attr(all(target_arch = "wasm32"), export_name = "free")]
#[no_mangle]
pub unsafe extern "C" fn _deallocate(ptr: u32, size: u32) {
    let layout =
        Layout::from_size_align(size as usize, std::mem::align_of::<u8>()).expect("Bad layout");
    dealloc(ptr as *mut u8, layout);
}

// Custom allocator implementation (for demonstration purposes, you can use `std::alloc::System` or other allocators).
#[global_allocator]
static ALLOC: System = System;
