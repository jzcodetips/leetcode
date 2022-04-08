use std::cmp;

fn main() {
    let s = 7;
    let arr = vec![2, 1, 5, 2, 8];

    println!("find_min_sub_array: {}", find_min_sub_array(s, &arr));
}

fn find_min_sub_array(s: i64, arr: &Vec<i64>) -> i64 {
    let mut smallest = i64::MAX;
    let mut sum = 0;
    let mut start = 0;

    for i in  0..arr.len() {
        sum += arr[i];

        while sum >= s {
            smallest = cmp::min(smallest, i as i64 - start + 1);
            sum -= arr[start as usize];
            start += 1;
        }
    }

    if smallest == i64::MAX {
        return 0;
    }

    smallest
}