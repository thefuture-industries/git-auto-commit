package constants

var EXPANDING_NOTATION_FOLDERS = map[string]string{
	"user":            "user management",
	"legacy":          "legacy code",
	"db":              "database layer",
	"tests":           "test suite",
	"payment":         "payment system",
	"api":             "API definitions", // OpenAPI/Swagger specs:contentReference[oaicite:0]{index=0}
	"auth":            "authentication module",
	"assets":          "static assets", // project assets (images, etc.):contentReference[oaicite:1]{index=1}
	"app":             "application code",
	"bin":             "executable scripts",
	"build":           "compiled outputs",   // compiled files:contentReference[oaicite:2]{index=2}
	"cmd":             "command-line tools", // main applications:contentReference[oaicite:3]{index=3}
	"client":          "client-side code",
	"common":          "common code",
	"config":          "configuration files", // config templates:contentReference[oaicite:4]{index=4}
	"controllers":     "controller logic",    // handles requests:contentReference[oaicite:5]{index=5}
	"core":            "core functionality",
	"data":            "data files",
	"dist":            "distribution packages", // compiled/distribution files:contentReference[oaicite:6]{index=6}
	"docs":            "documentation files",   // reference docs:contentReference[oaicite:7]{index=7}
	"examples":        "example applications",  // sample code:contentReference[oaicite:8]{index=8}
	"handlers":        "request handlers",
	"helpers":         "helper functions",
	"internal":        "internal packages", // private code:contentReference[oaicite:9]{index=9}
	"lib":             "library code",      // reusable libraries:contentReference[oaicite:10]{index=10}
	"log":             "log files",
	"middleware":      "middleware components",
	"models":          "data models",   // data schemas:contentReference[oaicite:11]{index=11}
	"public":          "public assets", // static files for web:contentReference[oaicite:12]{index=12}
	"resources":       "resource files",
	"routes":          "route handlers", // API endpoints:contentReference[oaicite:13]{index=13}
	"scripts":         "build scripts",  // automation scripts:contentReference[oaicite:14]{index=14}
	"services":        "service layer",
	"shared":          "shared code",
	"src":             "source code files", // main source directory:contentReference[oaicite:15]{index=15}
	"test":            "test suite",        // automated tests:contentReference[oaicite:16]{index=16}
	"tools":           "utility tools",     // development tools:contentReference[oaicite:17]{index=17}
	"vendor":          "vendor libraries",  // third-party dependencies:contentReference[oaicite:18]{index=18}
	"views":           "view templates",    // HTML/UI templates:contentReference[oaicite:19]{index=19}
	"tmp":             "temporary files",
	"pkg":             "public reusable packages",
	"third_party":     "third-party source code",
	"external":        "external dependencies",
	"fixtures":        "test fixtures and mock data",
	"mocks":           "mock implementations for testing",
	"stubs":           "stub implementations",
	"schemas":         "database or API schemas",
	"sql":             "SQL migration files",
	"migrations":      "database migration scripts",
	"seeds":           "database seed data",
	"proto":           "Protobuf definitions",
	"grpc":            "gRPC service definitions and stubs",
	"integration":     "integration tests",
	"functional":      "functional tests",
	"e2e":             "end-to-end tests",
	"ui":              "user interface components",
	"frontend":        "frontend source code",
	"backend":         "backend source code",
	"ops":             "operations scripts",
	"infra":           "infrastructure as code",
	"ci":              "continuous integration configs",
	"cd":              "continuous deployment configs",
	"deployment":      "deployment scripts and configs",
	"k8s":             "Kubernetes manifests",
	"helm":            "Helm charts",
	"ansible":         "Ansible playbooks",
	"terraform":       "Terraform configs",
	"docker":          "Dockerfiles and related configs",
	"monitoring":      "monitoring configuration and scripts",
	"logging":         "logging configuration",
	"analytics":       "analytics scripts or configs",
	"locales":         "localization and translation files",
	"i18n":            "internationalization files",
	"l10n":            "localization files",
	"themes":          "UI themes and styles",
	"styles":          "CSS/SCSS style files",
	"fonts":           "font files",
	"images":          "image resources",
	"audio":           "audio resources",
	"video":           "video resources",
	"reports":         "generated reports",
	"notebooks":       "Jupyter or data science notebooks",
	"research":        "research documents or code",
	"sandbox":         "experimental code",
	"playground":      "scratch code for testing ideas",
	"archive":         "archived old code",
	"deprecated":      "deprecated code",
	".github":         "GitHub configuration files",
	".husky":          "Git hooks managed by Husky",
	".vscode":         "VSCode editor settings",
	".idea":           "JetBrains IDE configuration",
	".config":         "User or app configuration files",
	".docker":         "Docker configuration files",
	".circleci":       "CircleCI configuration files",
	".gitlab-ci":      "GitLab CI/CD configuration files",
	".github_actions": "GitHub Actions workflows",
	".terraform":      "Terraform state and configs",
	".npm":            "npm package configs and cache",
	".cache":          "Cache files",
	".env":            "Environment variable files",
	".eslint":         "ESLint configuration files",
	".prettier":       "Prettier configuration files",
	".stylelint":      "Stylelint configuration files",
	".scripts":        "Custom shell or utility scripts",
}
