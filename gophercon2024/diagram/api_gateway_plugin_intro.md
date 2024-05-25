```mermaid
sequenceDiagram
  participant Client
  participant AG as API Gateway 
  participant AGP as API Gateway Plugin
  participant Server as Server
  Client ->> AG: Request
  AG->>AGP: Request
  AGP->>Server: Request
  Server->>AGP: Response
  AGP->>AG:Response
  AG->>Client:Response
```