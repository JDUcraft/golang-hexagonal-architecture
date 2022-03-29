# Hexagonal Architecture

## Principle

### Directories
* **domain**:
* **application**: 
* **infrastructure**: 

### Dependency


## Run

```bash
git clone git@github.com:JDUcraft/golang-hexagonal-architecture.git
cd golang-hexagonal-architecture
go run .
```

## Dependency graph
```mermaid
  graph TD;
      Application --> Domain;
      Infrastructure --> Domain;
```

```mermaid
gitGraph:
options
{
    "nodeSpacing": 150,
    "nodeRadius": 10
}
end
commit
branch newbranch
checkout newbranch
commit
commit
checkout master
commit
commit
merge newbranch
```
