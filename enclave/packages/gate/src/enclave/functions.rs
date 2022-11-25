use sgx_types::sgx_status_t;
use crate::enclave::consts::ENCLAVE_FILE_NAME;
use crate::enclave::enclave_api::ecall_health_check;
use crate::enclave::init::init_enclave;

pub fn health_check() -> Result<String, crate::Error> {

    let enclave = init_enclave(ENCLAVE_FILE_NAME).unwrap();

    let eid = enclave.geteid();
    let mut retval = sgx_status_t::SGX_SUCCESS;
    let status = unsafe {
        ecall_health_check(
            eid,
            &mut retval,
        )
    };

    if status != sgx_status_t::SGX_SUCCESS {
        println!("could not generate attestation report");
        panic!("omg");
    }

    if retval != sgx_status_t::SGX_SUCCESS {
        println!("could not generate attestation report");
        panic!("omg");
    }

    return Ok("success from enclave".to_string())
}
