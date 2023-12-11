# Code Organization

- Some codebases are easy to navigate.
- Others are hard to follow.
- Code organization is often the difference.

A good code structure will:
- Make it easier to guess where the bugs are.
- Make it easier to add a new feature.
- And more.

## Flat Structure

All code is in a single package.

Files are still separate code.

```bash
myapp/
    gallery_store.go
    gallery_handler.go
    gallery_templates.go
    user_store.go
    user_handler.go
    user_templates.go
    router.go

```

## Separation of concerns

Separate code based on duties.

HTMS & CSS is a traditional example:
- HTML focuses on overall structure.
- CSS focuses on syling it.

Model-View-Controller (MVC) is a ppopular structure following this strategy.
In a web app:
- `models` => data, logic, rulles; usually database.
- `views` => rendering things, usually HTML.
- `controller` => connects it all. Accepts user input, passes that to models to do stuff, 
then passes the data to views to render things; usually handlers.


```bash
myapp/
    controllers/
        user_handler.go
        gallery_handler.go
        ...
    views/
        user_templates.go
        ...
    models/
        user_store.go
        ...
```

Doesn't need to be named `models` `views` and `controllers`.
[Buffalo](https://gobuffalo.io/en/) uses a variation of MVC but doesn't name
the directories exactly.

Very common in framework b/c it is predictable. Makes it easier for a framework to guess where certain files will be.

Ruby on rails, Django, etc use MVC.

## Dependecy-based structure

Ben Johnson writes about this at [GO Beyound](https://www.gobeyond.dev/standard-package-layout/)

Structured based on dependency, but with a common set of interfaces and types they all work with.

```
myapp/
    user.go # defines a User type that isn't locked into any speciffic data store.
    use_store.go # defines UserStore interface.

    psql/
        user_store.go # Implements the UserStore interface with a Postgresql speciffic implementation.
```

## Many more...

- Domain driven design.
- Onion architecture.
- ...

## Which one should you use

No such thing as one-size fits all.
each has pros & cons.

Mat Ryer and David Hernandez used flat at <space.dev> early on (no idea now).
Flat can work well with microservices since each is small.
Flat won't work well at Google with a massive monorepo.

Experience also matters:
- Globals are bad, but why?
- Testable code only makes sense after you try to test untestable code.

My goal is to make this course accessible!

*Remember: You can always refactor.*

Don't get decision paralysis!

## We are going to use MVC

It isn't perfect, but it:
- Is predictable.
- Will provide us with some clear rules to follow when organizing our code.
- It will prevent us from getting stuck trying to decide how to organize our code.
- Is easy to pick up and utilize when learning. (Many others are not and are better suited
to experienced developers and teams).

Later we can explore how we might refactor to something different from MVC
*when we understand our needs better.*



