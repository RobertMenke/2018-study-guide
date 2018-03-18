use std::io::BufRead;
use std::io::stdin;
use std::iter::Iterator;

struct Instruction {
    instruction_type: u32,
    instruction_value: u32,
}

fn main() {
    let mut stack: Vec<u32> = vec![];
    let io_in = stdin();
    let lock = io_in.lock();
    for text in lock.lines() {
        let line = text.unwrap();
        let text_vec : Vec<&str> = line.split(" ").collect();
        let instruction = parse_line(text_vec);
        stack = match instruction.instruction_type {
            0 => stack,
            1 => {
                stack.push(instruction.instruction_value);
                stack
            },
            2 => {
                stack.pop();
                stack
            },
            3 => max_value(stack),
            _ => stack
        }
    }
}

fn parse_line(line: Vec<&str>) -> Instruction {
    let first = next_integer(&line, 0);
    let second = next_integer(&line, 1);

    return Instruction {
        instruction_type: if first > 3 { 0 } else { first },
        instruction_value: second,
    }
}

fn max_value(stack: Vec<u32>) -> Vec<u32> {
    let max = stack.iter().fold(0, |max, &current|
        if current > max { current } else { max },
    );

    //https://stackoverflow.com/questions/27734708/println-error-expected-a-literal-format-argument-must-be-a-string-literal
    println!("{}", max.to_string());

    return stack;
}

fn next_integer(line : &Vec<&str>, index : usize) -> u32 {
    line
        .get(index)
        .unwrap_or(&"0")
        .parse::<u32>()
        .unwrap()
}
