# Post-Quantum-Cryptography-Key-Pair-Gen

**Post-Quantum-Cryptography-Key-Pair-Gen** is an interface for generating key pairs based on modern post-quantum secure cryptographic algorithms. It uses **Dilithium** and **Kyber** ([Inspiration](https://github.com/g-utils/crystals-go)), both of which are among the leading algorithms in **Post-Quantum Cryptography** (PQC) that are considered resistant to quantum computer attacks.
<br>
<br>
Let's stay safe!

### Overview

In today's world, where quantum computers pose a potential threat to the security of classical encryption methods, **Post-Quantum Cryptography** has been developed to create algorithms that remain secure even in the age of quantum computing. This project uses **Dilithium** (for digital signatures) and **Kyber** (for encryption) to provide a secure mechanism for generating public and private key pairs.

### Features

- **Dilithium**: A modern, post-quantum secure digital signature algorithm known for its high efficiency and security.
- **Kyber**: A post-quantum secure encryption algorithm based on lattice problems, considered one of the most secure options against quantum attacks.
- Key pair generation where the **Dilithium** key comes first, followed by **Kyber**.
- **Signing** and **Verification** functions based on these algorithms.
- **Encryption** and **Decryption** functionality using Kyber for encryption and Dilithium for signing.

### Installation

Make sure to install the required libraries and dependencies:

```bash
go get github.com/g-utils/crystals-go/dilithium
go get github.com/g-utils/crystals-go/kyber
```

### Usage

**1. Generate Key Pair**
```go 
publicKey, privateKey := pqckpg_api.GenerateKeys(nil)
```

**2. Sign and verify a Message**
```go 
signMe := []byte("SIGN ME PLEASE!")
signature := pqckpg_api.Sign(decode(string(privateKey)), signMe)
isValid := pqckpg_api.Verify(decode(string(publicKey)), signMe, signature)
```

**3. Encrypt and Decrypt a Message**
```go 
var message string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ 12345678910"
encryptedMessage := pqckpg_api.Encrypt(decode(string(publicKey)), message)
decryptedMessage := pqckpg_api.Decrypt(decode(string(privateKey)), encryptedMessage)
```

### Contributing
I invite all developers and security experts to improve this project! If you have ideas on how to optimize or extend the system, please feel free to help!
<br>**Thanks for any support! - Stein**

### Bugs
If you find some. Fix them :)

### LICENSE
[LICENSE](./LICENSE)