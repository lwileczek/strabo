# Strabo
Document RESTful APIs which serve `application/json` data as fast and easy as possible. 
Strabo is not for documenting all endpoints and all use cases, but for very common and simple ones as quickly as possible.

### Idea
Take the data structures we use in our code already and have a tool read them and create documentation for them automatically. 
This project should quickly answer the questions: What and Where

   - What data do we need to send?
   - Where do we send it?
   - what data do we get back?

#### Guiding Principles
 - make as few changes as possible from the code used within a project to using Strabo to generate documentation
 - We return application
 - The documentation itself should be human readable

#### Assumptions
 - You're using RESTful APIs
 - You return `application/json`

## Getting Started
Please review the examples and syntax within the [documentation](./docs) directory.

## Status
> **Warning**
> This is in Alpha

Please checkout the [Road map](./ROADMAP.md)
