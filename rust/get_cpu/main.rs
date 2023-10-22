use num_cpus;

// get limit if pod
fn main() {
    let num_cores = num_cpus::get();
    println!("CPU Cores: {}", num_cores);
}