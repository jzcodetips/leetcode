use std::collections::HashMap;
use std::cmp;

fn main() {
    println!("longest_substring_k_distinct {}",
        longest_substring_k_distinct(&String::from("araaci"), 2));
    println!("longest_substring_k_distinct {}",
        longest_substring_k_distinct(&String::from("araaci"), 1));
    println!("longest_substring_k_distinct {}",
        longest_substring_k_distinct(&String::from("cbbebi"), 3));
    println!("longest_substring_k_distinct {}",
        longest_substring_k_distinct(&String::from("cbbebi"), 10));
}

fn longest_substring_k_distinct(s: &str, k: i64) -> i64 {
    let s_len = s.len() as i64;
    
    if k > s_len  {
        return s_len
    }

    let mut longest = 0;
    let mut letters: HashMap<char, i64> = HashMap::new();
    let mut start = 0;

    let str = s.as_bytes();

    for i  in 0..s_len {
        let i = i as usize;
        let c = str[i] as char;

        let counter = letters.entry(c).or_insert(0);
        *counter += 1;

        while letters.len() > k as usize {
            let c = str[start] as char;

            match letters.get_mut(&c) {
                Some(count) => {
                    if *count - 1 == 0 {
                        letters.remove(&c);
                    } else {
                        *count -= 1;
                    }
                },
                _ => (),
            }

            start += 1;
        }

        longest = cmp::max(longest, i - start + 1);
    }

    longest as i64
}