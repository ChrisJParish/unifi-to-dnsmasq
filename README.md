# Unifi to Dnsmasq

This Go program converts Unifi user data to a dnsmasq-hosts file.

## Code Overview

The code snippet provided creates a new file named "dnsmasq-hosts". If the file creation fails, the program will panic and terminate.

It then iterates over a list of users. For each user, it checks if the user has a fixed IP. If not, it skips to the next user.

The user's name is then converted to lowercase. If the user has a note, the note is used as the name. If not, all instances of a unsafe characters (:,space etc) in the name are replaced with a hyphen.

Finally, the user's fixed IP and name are written to the file. Each entry is written on a new line.

## Usage

To use this program, you need to have Go installed on your machine. Then, you can run the program with the command `go run unifi_to_dnsmasq.go -u <username> -h <unifi host> -p <unifi password>`.

Please ensure that you have the necessary permissions to create and write to files in the directory where you run the program.

## Contributing

Contributions are welcome. Please open an issue to discuss your ideas before making a pull request.

## License

This project is licensed under the MIT License.