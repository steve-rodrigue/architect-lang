# Architect Language

Architect is an AI-native software architecture language and compiler.

It describes domain models, services, APIs, event consumers, deployments, and versioned systems at a high level. The architecture becomes the source of truth, enabling deterministic code generation, refactoring, and long-term evolution of software systems.

> [!NOTE]
> Architect is under active development. The parser, source loading, and project model resolution are already implemented. The code generation pipeline and execution engine are currently being finalized, with the first end-to-end implementation expected at the beginning of next week.

![Architect Logo](https://raw.githubusercontent.com/steve-rodrigue/architect-lang/refs/heads/main/assets/logo.png)

---

# Table of Contents

- [Project Structure](#project-structure)
- [What Refers To What](#what-refers-to-what)
- [Project Definition](#project-definition)
  - [Project Options](#project-options)
- [Objects](#objects)
  - [Object Field Modifiers](#object-field-modifiers)
  - [Object Modifiers](#object-modifiers)
- [Common Types](#common-types)
  - [Built-in Value Types](#built-in-value-types)
  - [Number Constraints](#number-constraints)
- [Services](#services)
  - [Service Options](#service-options)
- [Applications](#applications)
  - [Application Options](#application-options)
- [Endpoints](#endpoints)
  - [Endpoint Options](#endpoint-options)
- [Consumers](#consumers)
- [Workflow Actions](#workflow-actions)
  - [fetch](#fetch)
  - [create](#create)
  - [store](#store)
  - [update](#update)
  - [emit](#emit)
  - [call](#call)
  - [return](#return)
- [Workflow Expressions](#workflow-expressions)
  - [Comparators](#comparators)
- [Deployments](#deployments)
  - [Deployment Options](#deployment-options)
- [Reference Summary](#reference-summary)
- [Goal](#goal)

---

# Project Structure

A project contains one folder per version. Each version contains shared objects, services, and deployments.

Applications live inside services. Endpoints and consumers live inside applications.

```text
stadan/
├── project.arch
│
├── 0.1.0/
│   ├── objects/
│   │   ├── User.arch
│   │   ├── Post.arch
│   │   ├── PostHistory.arch
│   │   └── PostEmbedding.arch
│   │
│   ├── services/
│   │   ├── maindb/
│   │   │   └── service.arch
│   │   └── blog_api/
│   │       ├── service.arch
│   │       └── application/
│   │           ├── application.arch
│   │           └── endpoints/
│   │               ├── CreatePost.arch
│   │               └── UpdatePost.arch
│   │
│   └── deployments/
│       ├── dev.arch
│       └── prod.arch
│
└── 0.1.1/
    ├── objects/
    ├── services/
    └── deployments/
```
---

# What Refers To What

Architect uses references between files and declarations.

```text
project.arch
└── references version folders and version files

version
├── references object files
├── references service files
└── references deployment files

service.arch
├── defines one runtime service
├── may reference one application file
├── may store object names
├── may declare event names
└── may explicitly depend on other services when no application is provided

application.arch
├── defines one application inside a service
├── references endpoint files
└── references consumer files

endpoint files
├── define HTTP API operations
├── use shared workflow actions
└── reference services by name when fetching, storing, emitting, or calling

consumer files
├── define event handlers
├── consume one event name
└── use shared workflow actions

deployment files
├── define one environment
├── reference service names
└── define runtime overrides for those services
```

---

# Project Definition

The root project.arch file describes available versions and the files included in each version.

```text
project Stadan {
  version "0.1.0" {
    services {
      "services/maindb/service.arch"
      "services/graphdb/service.arch"
      "services/blog_api/service.arch"
    }

    objects {
      "objects/User.arch"
      "objects/Post.arch"
      "objects/PostHistory.arch"
    }

    deployments {
      "deployments/dev.arch"
      "deployments/prod.arch"
    }

    next_version "0.1.1" {
      services {
        "services/maindb/service.arch"
      }

      deployments {
        "deployments/prod.arch"
      }
    }
  }
}
```

## Project Options

A project supports:

```text
project ProjectName {
  version "0.1.0" {
    services {
      "services/service_name/service.arch"
    }

    objects {
      "objects/ObjectName.arch"
    }

    deployments {
      "deployments/dev.arch"
    }

    next_version "0.1.1" {
      services {
      }

      objects {
      }

      deployments {
      }
    }
  }
}
```
---

# Objects

Objects define shared domain models for a project version.

```text
object User {
  id UUID primary
  email String unique
  displayName String?
}
```

History objects are supported:

```text
object PostHistory history_of Post {
  id UUID primary
  post Post
  changedAt Time
}
```

## Object Field Modifiers

Available field modifiers:

```text
primary
unique
renamed_from PreviousFieldName
deprecated
```

Example:

```text
object User {
  id UUID primary
  email String unique
  fullName String renamed_from displayName
  oldField String deprecated
}
```

## Object Modifiers

Available object modifiers:

```text
history_of ObjectName
renamed_from PreviousObjectName
deprecated
```

Example:

```text
object PostHistory history_of Post {
  id UUID primary
  post Post
}
```

---

# Common Types

A type reference has this shape:

```text
TypeName
TypeName?
TypeName+
TypeName*
TypeName[min,max]
TypeName[min,]
TypeName[,max]

List<TypeName>
List<TypeName>?
```

Examples:

```text
id UUID
name String
displayName String?
age Int[0,120]
score Float[0,1]
tags List<String>
posts List<Post>?
```

## Built-in Value Types

Current literal value kinds:

text string int float bool 

Examples:

architect name = "Steve" count = 10 score = 0.98 enabled = true 

## Number Constraints

Available number constraint kinds:

```text
none
one_or_more
zero_or_more
range
```

Syntax:

```text
Int+        // one_or_more
Int*        // zero_or_more
Int[0,10]   // range with min and max
Int[3,]     // range with min only
Int[,100]   // range with max only
```

---

# Services

A service describes a runtime node in the ecosystem.

Application-backed service:

```text
service BlogAPI go {
  exposes 8080
  application "application/application.arch"
}
``` 

Infrastructure service:

```text
service MainDB postgres {
  version "17"
  image "postgres:17"
  exposes 5432

  stores User
  stores Post
  stores PostHistory
}
```

Event bus service:

```text
service Events hatchet {
  version "latest"

  event PostCreated
  event PostUpdated
  event PostScored
}
```

## Service Options

Available service kinds:

```text
go
python
postgres
memgraph
milvus
hatchet
nginx
redis
prometheus
```

Available service properties:

```text
version "..."
image "..."
exposes 8080
depends_on ServiceName
application "application/application.arch"
stores ObjectName
event EventName
```

Dependency rule:

```text
If application is provided:
  dependencies should be inferred from the application workflow.

If application is not provided:
  dependencies should be declared explicitly using depends_on.
```

---

# Applications

An application lives inside a service and references endpoints and/or consumers.

API application:

```text
application BlogAPI {
  emits rest on 8080

  endpoints {
    "endpoints/CreatePost.arch"
    "endpoints/GetPost.arch"
  }

  consumers {
  }
}
```

Worker application:

```text
application DetectorWorker {
  listens events on 9090

  endpoints {
  }

  consumers {
    "consumers/PostCreated.arch"
    "consumers/PostUpdated.arch"
  }
}
```

## Application Options

Port directions:

```text
emits
listens 
```

Port kinds:

```text
rest 
events 
```

Application file blocks:

```text
endpoints {
  "endpoints/File.arch"
}

consumers {
  "consumers/File.arch"
}
```

---

# Endpoints

Endpoints define HTTP API behavior.

```text
endpoint CreatePost POST "/posts" {
  input {
    title String body
    content Text body
    postedBy UUID body
  }

  create post:Post {
    title = input.title
    content = input.content
    postedBy = input.postedBy
    createdAt = now()
  }

  store post to MainDB

  emit event:PostCreated {
    post = post
  }

  return post
}
```

## Endpoint Options

HTTP methods:

```text
GET 
POST 
PUT 
PATCH 
DELETE
``` 

Input source kinds:

```text
path 
query 
body 
header 
cookie 
``` 

Input source rules:

```text
id UUID path
id UUID body?
id UUID (path?|body?)
```

Meaning:

```text
path      required path input
body?     optional body input
(path?|body?) exactly one of path or body may provide the value
```

---

# Consumers

Consumers react to one event.

```text
consumes PostCreated {
  fetch post:Post from MainDB where id == event.post.id

  call EmbeddingService.Generate {
    text = post.content
  }

  return post
}
```

A consumer contains workflow actions, just like endpoints.

---

# Workflow Actions

Endpoints and consumers share the same workflow language.

Available actions:

```text
fetch
create
store
update
emit
call
return
```

## fetch

Fetches data from a service.

```text
fetch post:Post from MainDB where id == input.id
```

References:

```text
post      local variable name
Post      object/type name
MainDB    service name
input.id  endpoint input selector
```

## create

Creates a new object or event value.

```text
create post:Post {
  title = input.title
  content = input.content
}
```

## store

Stores a variable into a service.

```text
store post to MainDB
```

References:

```text
post    local variable
MainDB  service name
```

## update

Updates an existing variable.

```text
update post {
  title = input.title
  updatedAt = now()
}
```

## emit

Emits an event.

```text
emit event:PostCreated {
  post = post
}
```

References:

```text
event        local event variable
PostCreated  event type/name
```

## call

Calls another service or function.

```text
call EmbeddingService.Generate {
  text = post.content
}
```

References:

```text
EmbeddingService  service name
Generate          operation/function name
```

## return

Returns a workflow expression.

```text
return post
```

---

# Workflow Expressions

Available expression kinds:

```text
identifier
selector
literal
function_call
```

Examples:

```text
post
input.title
event.post.id
result.vector

now()
slugify(input.title)
calculateScore(input.title, 12, 12.5, true, false)
```

Special identifiers:

```text
input   endpoint input object
event   consumed or emitted event object
result  result from a previous call
```

## Comparators

Available comparators:

```text
==
!=
<
<=
>
>=
```

Example:

```text
fetch post:Post from MainDB where post.score >= 10
```

---

# Deployments

A deployment defines one runtime environment.

Development:

```text
deployment dev {
  vendor docker_compose

  service BlogAPI {
    replicas 1
  }

  service MainDB {
    volume "main-db-data"
  }
}
```

Production:

```text
deployment prod {
  vendor ansible
  inventory "inventories/prod.ini"

  service BlogAPI {
    replicas 3
    domain "api.stadan.org"
  }

  service MainDB {
    host "prod-db-01"
    volume "/var/lib/postgresql/data"
    backup true
  }
}
```

## Deployment Options

Available vendors:

```text
docker_compose
ansible
kubernetes
```

Deployment-level properties:

```text
vendor docker_compose
inventory "inventories/prod.ini"
```

Service deployment properties:

```text
service ServiceName {
  replicas 3
  domain "api.example.com"
  host "prod-db-01"
  volume "/var/lib/postgresql/data"
  backup true
}
```

backup supports:

```text
true
false
```

---

# Reference Summary

Relationship overview:

```text
Project
  references Version definitions

Version
  references Service files
  references Object files
  references Deployment files

Object
  references other Object names through field types

Service
  references Application file
  references Object names through stores
  references Event names through event
  references other Service names through depends_on

Application
  references Endpoint files
  references Consumer files

Endpoint
  references Object/type names
  references Service names through fetch/store/call
  emits Event names

Consumer
  references one Event name
  references Object/type names
  references Service names through fetch/store/call
  emits Event names

Deployment
  references Service names
```

---

# Goal

Architect aims to become the source of truth for software systems.

Instead of generating architecture from code, code is generated from architecture.

This enables:

```text
deterministic code generation
AI-assisted development
automated refactoring
versioned evolution
infrastructure generation
deployment generation
long-term maintainability
```