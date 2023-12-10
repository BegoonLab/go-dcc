# NRMA Digital Command Control (DCC) Implementation in Go

This project is a comprehensive kit for controlling model trains using the NMRA Digital Command Control (DCC) protocol. It includes software and hardware design, integration with Raspberry Pi, and a mobile-friendly application for seamless control on mobile devices. In addition, the project offers a solution for controlling train modules, such as servos and lights, in a model train diorama.

The implementation is based on the following NMRA standards:

* [S-9.1 Electrical Standard](https://www.nmra.org/sites/default/files/standards/sandrp/pdf/s-9.1_electrical_standards_for_digital_command_control_2021.pdf)
* [S-9.2 DCC Communications Standard](http://www.nmra.org/sites/default/files/s-92-2004-07.pdf)
* [S-9.2.1 Extended Packet Formats for Digital Command Control Standard](http://www.nmra.org/sites/default/files/s-9.2.1_2012_07.pdf)

## Features

- Full DCC protocol implementation for controlling model trains and DCC Accessories modules
- Hardware design for Raspberry Pi integration
- Mobile-friendly application with zero overhead, compatible with any mobile phone
- Control solution for train modules, including servos and lights in model train dioramas

## Demo
[Mobile client application demo](https://begoonlab.tech/demos/train-commander/demo.index.html) 

## Installation

1. Clone this repository:
```bash
git clone https://github.com/yourusername/yourrepository.git
```

## Usage
Ensure your Raspberry Pi is set up and connected to your model train and DCC Accessories modules.

Run the compiled executable to start the server:

```bash
./yourrepository
```
Open the mobile-friendly application on your phone by navigating to the server's IP address and port number.

Use the application's interface to control your model train and DCC Accessories modules.

## Contributing
We welcome contributions to this project! Please follow these steps:

1. Fork the repository
2. Create a new branch (`git checkout -b your-feature`)
3. Commit your changes (`git commit -m 'Add a new feature`)
4. Push to the branch (`git push origin your-feature`)
5. Create a new Pull Request


## License
This project is licensed under the MIT License.