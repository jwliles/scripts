use env_logger;
use log::{error, info};

pub fn init_logger() {
    env_logger::init();
}

pub fn _log_event(message: &str) {
    info!("{}", message);
}

pub fn _log_error(message: &str) {
    error!("{}", message);
}
