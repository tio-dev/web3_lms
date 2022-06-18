# web3_lms
Sistema de Gestão de Aprendizagem (LMS - Learning Management System) de código aberto utilizando tecnologias da WEB3

É uma plataforma de aprendizagem que será construída no curso "WEB3 para Patos" como prática do aprendizado, porém extrapolando os fins didáticos e podendo ser utilizada em produção livremente.

A pilha de tecnologias é opinativa. São as tecnologias, ferramentas e práticas utilizadas pelo nosso time na empresa STEMativa (https://stemativa.com.br), especializada em soluções descentralizadas para o mercado agro, florestal e de capitais.

## Pilha tecnológica

**Monorepo com NX**

Autenticação e autorização com KeyCloak
- OpenID Connect 
- Papeis e permissões baseados em grupos de usuário
- [Authentication and authorization using the keycloak REST API](https://developers.redhat.com/blog/2020/11/24/authentication-and-authorization-using-the-keycloak-rest-api#keycloak_connection_using_a_java_application) 

**FrontEnd com TypeScript**
- FrontEnd com React/Typescript/NextJS
- Controle de estado da aplicação com MobX
- Base local (no browser) com PouchDB sincronizada com bases remotas em CouchDB
- UI Kit / framework CSS com MUI (Material UI)
- Cliente GraphQL com graphql-hooks e relay
- Deploy na Vercel

**Mobile com Flutter**
- Controle de estado da aplicação com MobX
- Base de dados local com PouchDB sincronizada com bases remotas em CouchDB

**BackEnd com GO**
- Framework WEB com Gin Gonic
- API GraphQL com gqlgqn (Mutations)
- ORM com gorm e smallnest/gen
- Lógica e responsabilidades distribuída em pequenos serviços assíncronos (ver tópico de mensaggeria)

Banco de dados com PostgreSQL
- HeadlessCMS com Directus
- Hub Graphql com Hasura
- Base de dados analítica segregada da transacional
- GraphQL (Query e Subscriptions) na base analítica
- Denormalização e uso de JSON documents para performance

**Mensageria com NATS**
- Comunicação baseada em eventos e abordagem pub/sub
- Stream de mensagens com JetStream

**Deploy com HashiCorp e Docker**
- Receitas Terraform
- Receitas Waypoint

**Blockchain Algorand**
- Carteira MOBILE com Pera / Algorand Wallet
- Integração com carteiras com WalletConnect protocol
- Aceleração e simplificação de desenvolvimento com API e SDK do Boa-Fé Community Edition
