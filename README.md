### **Monitor Disk Usage**

![Static Badge](https://img.shields.io/badge/CheStix-PassGenerator-PassGenerator)
![GitHub top language](https://img.shields.io/github/languages/top/CheStix/PassGenerator)
![GitHub](https://img.shields.io/github/license/CheStix/PassGenerator)
![GitHub Repo stars](https://img.shields.io/github/stars/CheStix/PassGenerator)
![GitHub issues](https://img.shields.io/github/issues/CheStix/PassGenerator)
- **Description**: This is a password generator CLI application that not only generates passwords, but also encrypts and stores them - so it's really functional.
- **Key Features**:
    - AES symmetric key encryption algorithm for encrypting and decrypting your password.
    - Passwords are stored in a file in encrypted binary form. And if the file is ever exposed, your passwords will be safe because they will be encrypted and really hard to read.
    - The ability to set the complexity for the generated password.
- **Usage**:
    -  Get the executable file:
       - download the source code and run `go build PassGenerator`
       - or just download the executable file [from here](https://github.com/CheStix/PassGenerator/releases/download/v1.0.0/PassGenerator.exe )
  - Run with the parameters **"--get"** to get the password and **"--set"** to generate and save the password. When generating a password, the **platform** parameter is specified with reference to which the encrypted password will be stored, the length of the password **len** and the complexity **level**:
     - basic generates random strings of small letters.
     - medium takes some of the characters from the basic function and adds random capital letters to it.
     - hard does the same for medium, but adds numbers to it.
     - xhard does the same thing and adds special characters.
- **Examples**:
  - ```cmd
    PassGenerator.exe --set platform=github len=8 level=xhard
    ```
  - ```cmd
    PassGenerator.exe --get platform=github
    ```