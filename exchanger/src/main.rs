use std::fmt;
use std::io;
use Temperature::*;

#[derive(Debug, PartialEq, Copy, Clone)]

///An enum representing the different Temperature scales.
///These are the available temperature scales to be converted between with this script.

pub enum Temperature {
    Kelvin(f64),
    Celsius(f64),
    Fahrenheit(f64),
    Rankine(f64),
}

impl fmt::Display for Temperature {
    fn fmt(&self, fmtr: &mut fmt::Formatter) -> fmt::Result {
        match *self {
            Kelvin(k)     => write!(fmtr, "{}K", k),
            Celsius(c)    => write!(fmtr, "{}°C", c),
            Fahrenheit(f) => write!(fmtr, "{}°F", f),
            Rankine(r)    => write!(fmtr, "{}°R", r),
        }
    }
}

impl Temperature {
    /// Convert the Temperature unit into Celsius
    pub fn to_celsius(self) -> Temperature {
        match self {
            Kelvin(k)      => Celsius(k - 273.15),
            c @ Celsius(_) => c,
            Fahrenheit(f)  => Celsius((f - 32.0) * (5.0 / 9.0)),
            Rankine(r)     => Celsius((r - 491.67) * 5.0 / 9.0),
        }
    }

    /// Convert the Temperature unit into Fahrenheit
    pub fn to_fahrenheit(self) -> Temperature {
        match self {
            Kelvin(k)         => Fahrenheit((k * (9.0 / 5.0)) - 459.67),
            Celsius(c)        => Fahrenheit((c * (9.0 / 5.0)) + 32.0),
            f @ Fahrenheit(_) => f,
            Rankine(r)        => Fahrenheit(r - 459.67),
        }
    }

    /// Convert the Temperature unit into Kelvin
    pub fn to_kelvin(self) -> Temperature {
        match self {
            k @ Kelvin(_) => k,
            Celsius(c)    => Kelvin(c + 273.15),
            Fahrenheit(f) => Kelvin((f + 459.67) * (5.0 / 9.0)),
            Rankine(r)    => Kelvin(r * 5.0 / 9.0),
        }
    }

    /// Convert the Temperature unit into Rankine
    pub fn to_rankine(self) -> Temperature {
        match self {
            r @ Rankine(_) => r,
            Celsius(c)     => Rankine((c + 273.15) * 9.0 / 5.0),
            Fahrenheit(f)  => Rankine(f + 459.67),
            Kelvin(k)      => Rankine(k * 9.0 / 5.0),
        }
    }
}

/// # Temperature Conversion Formulas.
/// These are the formulas used in the crate for temperature conversions.
///
/// | From       | To         | Formula                      | Notes |
/// | ----       | ---        | -------                      | ---   |
/// | Celsius    | Fahrenheit | F = (C * 9.0 / 5.0) + 32.0   |       |
/// | Celsius    | Kelvin     | K = C + 273.15               |       |
/// | Celsius    | Rankine    | R = (C + 273.15) * 9.0 / 5.0 |       |
/// | Fahrenheit | Celsius    | C = (F - 32.0) * 5.0 / 9.0   |       |
/// | Fahrenheit | Kelvin     | K = (F + 459.67) * 5.0 / 9.0 |       |
/// | Fahrenheit | Rankine    | R = F + 459.67               |       |
/// | Kelvin     | Celsius    | C = K - 273.15               |       |
/// | Kelvin     | Fahrenheit | F = K * 9.0 / 5.0 - 459.67   |       |
/// | Kelvin     | Rankine    | R = K * 9.0 / 5.0            |       |
/// | Rankine    | Celsius    | C = (R - 491.67) * 5.0 / 9.0 |       |
/// | Rankine    | Fahrenheit | F = R - 459.67               |       |
/// | Rankine    | Kelvin     | K = R * 5.0 / 9.0            |       |
///
/// # Memo:
/// The only temperature where the Kelvin and Fahrenheit values are equal is at 574.59.
/// The only temperature where the Celsius and Fahrenheit values are equal is at -40.

pub fn convert_temp(temperature: &Temperature, target_scale: &str) -> String {
    match *temperature {
        Temperature::Celsius(degrees) => match target_scale {
            "F" => format!("{}°C = {}°F", degrees, (degrees * 9.0 / 5.0) + 32.0),
            "K" => format!("{}°C = {}K", degrees, degrees + 273.15),
            "R" => format!("{}°C = {}°R", degrees, (degrees + 273.15) * 9.0 / 5.0),
            _ => String::from("Unsupported conversion"),
        },
        Temperature::Fahrenheit(degrees) => match target_scale {
            "C" => format!("{}°F = {}°C", degrees, (degrees - 32.0) * (5.0 / 9.0)),
            "K" => format!("{}°F = {}K", degrees, (degrees + 459.67) * (5.0 / 9.0)),
            "R" => format!("{}°F = {}°R", degrees, degrees + 459.67),
            _ => String::from("Unsupported conversion"),
        },
        Temperature::Kelvin(degrees) => match target_scale {
            "C" => format!("{}K = {}°C", degrees, degrees - 273.15),
            "F" => format!("{}K = {}°F", degrees, degrees * (9.0 / 5.0) - 459.67),
            "R" => format!("{}K = {}°R", degrees, degrees * (9.0 / 5.0)),
            _ => String::from("Unsupported conversion"),
        },
        Temperature::Rankine(degrees) => match target_scale {
            "C" => format!("{}°R = {}°C", degrees, (degrees - 491.67) * (5.0 / 9.0)),
            "F" => format!("{}°R = {}°F", degrees, degrees - 459.67),
            "K" => format!("{}°R = {}K", degrees, degrees * (5.0 / 9.0)),
            _ => String::from("Unsupported conversion"),
        },
    }
}

/// # Get the user input
/// Add a question to allow the user to specify which scale the temperature should be converted
/// into.

pub fn get_user_temp() {
    println!("\nType \"exit\" to exit the program");

    loop {
        let mut temp_input = String::new();

        println!(
            "\nPlease input a temperature you want to convert (Example: 10R, 10F, 10K, or \
        -10C):"
        );

        io::stdin()
            .read_line(&mut temp_input)
            .expect("Failed to read line");

        let trimmed = temp_input.trim();

        if trimmed == "exit" {
            break;
        }

        let (temp_str, scale_str) = trimmed.split_at(trimmed.len() - 1);

        let temp: f64 = match temp_str.parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        let temperature: Temperature = match scale_str.to_uppercase().as_str() {
            "C" => Temperature::Celsius(temp),
            "F" => Temperature::Fahrenheit(temp),
            "K" => Temperature::Kelvin(temp),
            "R" => Temperature::Rankine(temp),
            _ => continue,
        };

        let mut target_scales: Vec<Temperature> = Vec::new();
        println!("Enter the target scale(s) you want to convert to (CFKR), e.g., 'CF':");
        let mut target_input = String::new();
        io::stdin()
            .read_line(&mut target_input)
            .expect("Failed to read line");

        let target_input = target_input.trim().to_uppercase();

        for c in target_input.chars() {
            match c {
                'C' => target_scales.push(Temperature::Celsius(0.0)),
                'F' => target_scales.push(Temperature::Fahrenheit(0.0)),
                'K' => target_scales.push(Temperature::Kelvin(0.0)),
                'R' => target_scales.push(Temperature::Rankine(0.0)),
                _ => continue,
            }
        }

        println!("Converting {} to target scale {:?}", temperature, target_scales);
        println!("Result: {}", convert_temp(&temperature, &target_input));
    }
}


/// # Print the final converted temp scale
/// This will print the converted temperature into the temperature scale the user has requested.

pub fn print_temp(temperature: &Temperature, target_scales: &[Temperature]) {
    for target_scale in target_scales {
        println!("{}", convert_temp(temperature, &target_scale.to_string()));
    }
}

fn main() {
    println!("This is a temperature converter!\n");
    let target_scales: Vec<String> = vec!["C".to_string(), "F".to_string(), "K".to_string(), "R".to_string()];

    get_user_temp();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_to_celsius() {
        assert_eq!(Temperature::Kelvin(0.0).to_celsius(), Temperature::Celsius(-273.15));
        assert_eq!(Temperature::Fahrenheit(32.0).to_celsius(), Temperature::Celsius(0.0));
        // Add more test cases for edge conditions
        assert_eq!(Temperature::Kelvin(f64::MAX).to_celsius(), Temperature::Celsius(f64::MAX - 273.15));
    }

    #[test]
    fn test_to_fahrenheit() {
        if let Temperature::Fahrenheit(value) = Temperature::Celsius(0.0).to_fahrenheit() {
            assert!(f64::abs(value - 32.0) < 0.000001); // Adjust epsilon as needed
        } else {
            panic!("Conversion to Fahrenheit failed");
        }
    }

    #[test]
    fn test_to_kelvin() {
        assert_eq!(Temperature::Celsius(0.0).to_kelvin(), Temperature::Kelvin(273.15));
        assert_eq!(Temperature::Fahrenheit(32.0).to_kelvin(), Temperature::Kelvin(273.15));
        // Add more test cases as needed
    }

    #[test]
    fn test_to_rankine() {
        assert_eq!(Temperature::Celsius(0.0).to_rankine(), Temperature::Rankine(491.67));
        assert_eq!(Temperature::Fahrenheit(32.0).to_rankine(), Temperature::Rankine(491.67));
        // Add more test cases as needed
    }

    #[test]
    fn test_case_insensitivity() {
        let actual = convert_temp(&Temperature::Celsius(100.0), "F");
        let expected = "100.0°C = 212.0°F".to_uppercase(); // Convert both strings to uppercase
        assert_eq!(actual.to_uppercase(), expected);
    }

    #[test]
    fn test_convert_temp() {
        let actual = convert_temp(&Temperature::Celsius(0.0), "F");
        let expected = "0°C = 32°F".to_string();
        assert_eq!(actual, expected);
        // Add more test cases for error conditions and edge cases
    }

    #[test]
    fn test_display_format() {
        let temperature = Temperature::Celsius(100.0);
        let actual = format!("{:.1}", temperature); // Adjust format to match expected output
        let expected = "100.0°C".to_string();
        assert_eq!(actual, expected);
    }

    // Write more tests to cover other aspects of your code
}
