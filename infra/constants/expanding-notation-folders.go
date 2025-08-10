package constants

var EXPANDING_NOTATION_FOLDERS = map[string]string{
	"user":        "user management",
	"legacy":      "legacy code",
	"db":          "database layer",
	"tests":       "test suite",
	"payment":     "payment system",
	"api":         "API definitions", // OpenAPI/Swagger specs:contentReference[oaicite:0]{index=0}
	"auth":        "authentication module",
	"assets":      "static assets", // project assets (images, etc.):contentReference[oaicite:1]{index=1}
	"app":         "application code",
	"bin":         "executable scripts",
	"build":       "compiled outputs",   // compiled files:contentReference[oaicite:2]{index=2}
	"cmd":         "command-line tools", // main applications:contentReference[oaicite:3]{index=3}
	"client":      "client-side code",
	"common":      "common code",
	"config":      "configuration files", // config templates:contentReference[oaicite:4]{index=4}
	"controllers": "controller logic",    // handles requests:contentReference[oaicite:5]{index=5}
	"core":        "core functionality",
	"data":        "data files",
	"dist":        "distribution packages", // compiled/distribution files:contentReference[oaicite:6]{index=6}
	"docs":        "documentation files",   // reference docs:contentReference[oaicite:7]{index=7}
	"examples":    "example applications",  // sample code:contentReference[oaicite:8]{index=8}
	"handlers":    "request handlers",
	"helpers":     "helper functions",
	"internal":    "internal packages", // private code:contentReference[oaicite:9]{index=9}
	"lib":         "library code",      // reusable libraries:contentReference[oaicite:10]{index=10}
	"log":         "log files",
	"middleware":  "middleware components",
	"models":      "data models",   // data schemas:contentReference[oaicite:11]{index=11}
	"public":      "public assets", // static files for web:contentReference[oaicite:12]{index=12}
	"resources":   "resource files",
	"routes":      "route handlers", // API endpoints:contentReference[oaicite:13]{index=13}
	"scripts":     "build scripts",  // automation scripts:contentReference[oaicite:14]{index=14}
	"services":    "service layer",
	"shared":      "shared code",
	"src":         "source code files", // main source directory:contentReference[oaicite:15]{index=15}
	"test":        "test suite",        // automated tests:contentReference[oaicite:16]{index=16}
	"tools":       "utility tools",     // development tools:contentReference[oaicite:17]{index=17}
	"vendor":      "vendor libraries",  // third-party dependencies:contentReference[oaicite:18]{index=18}
	"views":       "view templates",    // HTML/UI templates:contentReference[oaicite:19]{index=19}
	"tmp":         "temporary files",
}
