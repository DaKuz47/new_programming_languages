use std::fs::File;
use std::io::{Write, Read};

fn main() {
    let args: Vec<_> = std::env::args().collect();
    
    let action = args[1].clone();
    let value = args[2].clone();
    let path = if args.len() == 4 {args[3].clone()} else {"todo.txt".to_owned()};

    let path_exists = std::path::Path::new(&path).exists();

    if action == "add" {
        let mut file = std::fs::OpenOptions::new().create(true).append(true).open(&path);
        let ln_value = value.clone() + "\n";
        file.expect("REASON").write(ln_value.as_bytes());
    }

    if action == "del" {
        let mut file = File::open(path.clone()).unwrap();
        let mut file_contents = String::new();
        file.read_to_string(&mut file_contents).unwrap();

        let mut lines = file_contents.split("\n").collect::<Vec<&str>>();
        let line_to_remove: usize = value.parse().unwrap();

        lines.remove(line_to_remove);
        let joined = lines.join("\n");

        std::fs::remove_file(&path);
        let mut file = std::fs::OpenOptions::new().create(true).append(true).open(&path);
        file.expect("REASON").write(joined.as_bytes());
    }
}
