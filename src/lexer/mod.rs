#[cfg(test)]
mod test;

#[allow(dead_code)]
#[derive(Debug, PartialEq)]
pub enum Token {
    Ident(String),
    Int(String),

    Illegal,
    Eof,
    Assign,

    Plus,
    Minus,
    ForwardSlash,
    Asterisk,
    Equal,
    NotEqual,
    LessThan,
    GreaterThan,

    Not,
    Comma,
    Semicolon,
    Lparen,
    Rparen,
    Lbracket,
    Rbracket,

    Let,
    Const,
    Var,

    Func,
    If,
    Else,
    Return,
    True,
    False,
}

#[derive(Debug)]
pub struct Lexer {
    position: usize,
    read_position: usize,
    current_character: u8,
    input: Vec<u8>,
}
