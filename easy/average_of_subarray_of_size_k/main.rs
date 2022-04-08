/*
    Given an array, find the average of all subarrays of ‘K’ contiguous elements in it.
*/

fn main() {
    let k = 2;
    let arr = vec![1, 3, 2, 6, -1, 4, 1, 8, 2];

    println!("find_averages_brut_force: {:?}", find_averages_brut_force(k, &arr));
    println!("find_averages_sliding_window: {:?}", find_averages_sliding_window(k, &arr));
}

// brut force
fn find_averages_brut_force(mut k: usize, arr: &Vec<i64>) -> Vec<f64> {
    let mut result: Vec<f64> = Vec::new();
    let len = arr.len();

    if k > len {
        k = len
    }

    for i in 0..arr.len()-k+1 {
        let mut sum = 0.0;

        for j in i..i+k {
            sum += arr[j] as f64;
        }

        result.push(sum/k as f64)
    }

    result
}

// sliding_window
fn find_averages_sliding_window(mut k: usize, arr: &Vec<i64>) -> Vec<f64> {
    let mut result: Vec<f64> = Vec::new();
    let len = arr.len();

    if k > len {
        k = len
    }

    let mut sum = 0.0;
    let mut start = 0;

    for i in 0..arr.len() {
        sum += arr[i] as f64;

        if i >= k - 1 {
            result.push(sum/k as f64);
            sum -= arr[start] as f64;
            start += 1;
        }
    }

    result
}