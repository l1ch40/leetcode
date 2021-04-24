use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut m = HashMap::new();
        let mut result: Vec<i32> = Vec::new();
        for (index, value) in nums.iter().enumerate() {
            let another = target - value;
            match m.get(&another) {
                Some(&inx) => {
                    result.push(inx);
                    result.push(index as i32);
                    return result;
                }
                None => (),
            }
            m.insert(value, index as i32);
        }
        return result;
    }
}
