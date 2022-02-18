
fn reverse1(s: &str) -> String {
    let mut ret = String::new();
    let cs: Vec<char> = s.chars().collect();
    let l = cs.len();
    for i in 0..l {
        ret += &cs[l-i-1].to_string();
    }
    return ret;
}

fn reverse2(s: &str) -> String {
    let mut s: Vec<char> = s.chars().collect();
    let l = s.len();
    for i in 0..l/2 {
        let tmp = s[i];
        s[i] = s[l-i-1];
        s[l-i-1] = tmp;
    }
    return s.into_iter().collect()
}

fn main() {
    let input = "kadokawa";
    println!("input: {}", input);

    println!("rev_1: {}", reverse1(&input));
    println!("rev_2: {}", reverse2(&input));
}

