use std::time::{Duration, Instant};

pub struct Metrics {
    pub files_scanned: u32,
    _start_time: Option<Instant>,
    elapsed_time: Option<Duration>,
}

impl Metrics {
    pub fn new() -> Self {
        Metrics {
            files_scanned: 0,
            _start_time: None,
            elapsed_time: None,
        }
    }

    pub fn _start_timer(&mut self) {
        self._start_time = Some(Instant::now());
    }

    pub fn _stop_timer(&mut self) {
        if let Some(_start_time) = self._start_time {
            self.elapsed_time = Some(_start_time.elapsed());
        }
    }

    pub fn increment_files_scanned(&mut self) {
        self.files_scanned += 1;
    }

    pub fn display_metrics(&self) {
        println!("Scan Statistics:");
        println!("Total files scanned: {}", self.files_scanned);
        if let Some(elapsed) = self.elapsed_time {
            self.display_time_metrics(elapsed);
        }
    }

    fn display_time_metrics(&self, elapsed: Duration) {
        println!("Total time taken: {:.2} seconds", elapsed.as_secs_f64());
        let rate = self.files_scanned as f64 / elapsed.as_secs_f64();
        println!("Average scan rate: {:.2} files/second", rate);
    }
}
