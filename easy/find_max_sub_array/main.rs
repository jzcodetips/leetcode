fn main() {
    let k = 3;
    let arr = vec![1, 2, 3, 4, 5, 6];

    println!("find_max_sub_array: {}", find_max_sub_array(k, &arr));
}

fn find_max_sub_array(k: usize, arr: &Vec<i64>) -> i64 {
    let mut max = 0;
    let mut curr_max = 0;
    let mut start = 0;

    for i in 0..arr.len() {
        curr_max += arr[i];

        if i >= k - 1 {
            if curr_max > max {
                max = curr_max
            }

            curr_max -= arr[start];
            start += 1;
        }
    }

    max
}
