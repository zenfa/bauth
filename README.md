# Background

## What is OAuth
From https://en.wikipedia.org/wiki/OAuth, "OAuth is an open standard for access delegation, commonly used as a way for Internet users to grant websites or applications access to their information on other websites but without giving them the passwords."

Basically, users can grant a client application to acquire an access token (which represents a user’s permission for the client to access their data) which can be used to authenticate a request to an API endpoint.

## Problems with OAuth
### Process is Tedious
from https://tools.ietf.org/html/rfc6749. Users are redirected back and forth to complete OAuth login.

     +--------+                               +---------------+
     |        |--(A)- Authorization Request ->|   Resource    |
     |        |                               |     Owner     |
     |        |<-(B)-- Authorization Grant ---|               |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(C)-- Authorization Grant -->| Authorization |
     | Client |                               |     Server    |
     |        |<-(D)----- Access Token -------|               |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(E)----- Access Token ------>|    Resource   |
     |        |                               |     Server    |
     |        |<-(F)--- Protected Resource ---|               |
     +--------+                               +---------------+
     
### Security Risks
Access tokens must be kept confidential in transit and in storage. Because anyone with the token can access the resource. 

# Solution - Blockchain Auth (BAuth)
Blockchain is ideal to provide identity service, thanks to public key cryptography and digital signature technology in blockchain.

Users can create Access Token with all necessary information(e.g. Client information, Scope, and Expiration Time) by themselves, and hand it over to clients. After that, clients can build requests and send to Blockchain network, attaching the Access Token. Dapp can verify the signature in the Access Token, and take actions only specified in it. 

     +---------------------------------------------+  
     |                                             |  
     | Access Token Created by Resource Owner      | 
     | (including Client, Scope, Signature, etc.)  |  
     +---------------------------------------------+
     |                                             |
     | Message Sent by Client                      |
     |                                             |  
     +---------------------------------------------+ 

## Simplified Process

     +--------+                               +---------------+
     |        |                               |   Resource    |
     |        |                               |     Owner     |
     |        |<------ Access Token ----------|               |
     | Client |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |---------- Access Token ------>|    Resource   |
     |        |                               |     Server    |
     |        |<------- Protected Resource ---|               |
     +--------+                               +---------------+

## More Secure, Under Control
In comparison to OAuth, the Access Token in BAuth defines who can use this token. 

Future Work: 
* to define signature verification in CustomSigVerify, with gas consumption
* to add timestamp and access token validation check

# Commands and Console Logs
```
cosmos@cosmoss-MacBook-Pro bauth % bauthcli tx bauth get-token $(bauthcli keys show agent -a) bank  --from user1
cosmos@cosmoss-MacBook-Pro bauth % bauthcli tx bauth access-resource $(bauthcli keys show user1 -a) bank 10token --from agent
{
  "chain_id": "bauth",
  "account_number": "3",
  "sequence": "0",
  "fee": {
    "amount": [],
    "gas": "200000"
  },
  "msgs": [
    {
      "type": "bauth/AccessResource",
      "value": {
        "owner": "cosmos14cyx0ps9ylfxjznh2v73nz90es4va80ajqzqzv",
        "client": "cosmos1v0esyjg8yhauaxwk2fqpxysuej87cz69eynxg8",
        "action": "bank",
        "amount": [
          {
            "denom": "token",
            "amount": "10"
          }
        ],
        "sig": "GUL7DOaa5reRijCzdORjN3ptnUusch+EZhnK4d66YvEvm+O3isu0l+BAOLBGLIAmdW90S60mICwQm3fMxEPXyg=="
      }
    }
  ],
  "memo": ""
}

confirm transaction before signing and broadcasting [y/N]: y
{
  "height": "0",
  "txhash": "3F0F5D4B56967CA9579116F785794EAADA8D5180648DB18FB549EE7992159DEA",
  "raw_log": "[]"
}
cosmos@cosmoss-MacBook-Pro bauth % bauthcli q account $(bauthcli keys show user1 -a)
{
  "type": "cosmos-sdk/Account",
  "value": {
    "address": "cosmos14cyx0ps9ylfxjznh2v73nz90es4va80ajqzqzv",
    "coins": [
      {
        "denom": "token",
        "amount": "990"
      }
    ],
    "public_key": {
      "type": "tendermint/PubKeySecp256k1",
      "value": "AuZQ00T8Kg5lDEYRNIfLt+MumGeTUtrWTX6s24FK489u"
    },
    "account_number": "2",
    "sequence": "1"
  }
}
cosmos@cosmoss-MacBook-Pro bauth % cat accessToken.txt
B�
  �淑�0�t�c7zm�K�r�f��޺b�/�㷊˴��@8�F,�&uotK�& ,�w��C��%                                                                                
```
