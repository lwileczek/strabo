use super::{Lexer, Token};

type DynResult<T> = Result<T, Box<dyn std::error::Error>>;

#[test]
fn get_next_token() -> DynResult<()> {
    let input = "=+(){},;";
    let mut lexer = Lexer::new(input.into());

    let tokens = vec![
        Token::Assign,
        Token::Plus,
        Token::Lparen,
        Token::Rparen,
        Token::Lbracket,
        Token::Rbracket,
        Token::Comma,
        Token::Semicolon,
    ];

    for token in tokens {
        let next_token = lexer.next_token()?;
        println!("expected: {:?}, received {:?}", token, next_token);
        assert_eq!(token, next_token);
    }

    Ok(())
}
