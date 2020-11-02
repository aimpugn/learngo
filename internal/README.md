# `/internal`
- [Source of Contents](https://github.com/golang-standards/project-layout/blob/master/internal)
## 설명
- Private 애플리케이셔 또는 라이브러리 코드 
- 다른 프로그램의 애플리케이션이나 라이브러리에서 임포트 하지 않길 원하는 코드
- Go 컴파일러 자체에 의해 강제되는 레이아웃 패턴임에 주의. 자세한 내용은 Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) 참조.
- 최상위 레벨의 `internal` 디렉토리에 제한되지 않는다는 점에 주의. 프로젝트 트리의 어떤 레벨에서든지 하나 이상의 `internal` 디렉토리를 가질 수 있다.

- You can optionally add a bit of extra structure to your internal packages to separate your shared and non-shared internal code. 
- It's not required (especially for smaller projects), but it's nice to have visual clues showing the intended package use. 
- Your actual application code can go in the `/internal/app` directory (e.g., `/internal/app/myapp`) and the code shared by those apps in the `/internal/pkg` directory (e.g., `/internal/pkg/myprivlib`).

Examples:

* https://github.com/hashicorp/terraform/tree/master/internal
* https://github.com/influxdata/influxdb/tree/master/internal
* https://github.com/perkeep/perkeep/tree/master/internal
* https://github.com/jaegertracing/jaeger/tree/master/internal
* https://github.com/moby/moby/tree/master/internal
* https://github.com/satellity/satellity/tree/master/internal

## `/internal/pkg`

Examples:

* https://github.com/hashicorp/waypoint/tree/main/internal/pkg