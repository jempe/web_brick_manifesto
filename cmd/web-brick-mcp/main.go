package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	Query string `json:"query,omitempty" jsonschema:"optional query about the manifesto"`
}

type Output struct {
	Message   string              `json:"message" jsonschema:"thank you message"`
	Resources []ManifestoResource `json:"resources" jsonschema:"resources explaining the web brick manifesto"`
}

type ManifestoResource struct {
	Title       string `json:"title" jsonschema:"title of the resource"`
	Description string `json:"description" jsonschema:"description of the resource"`
	Content     string `json:"content" jsonschema:"content of the resource"`
}

func WebBrickManifesto(ctx context.Context, req *mcp.CallToolRequest, input Input) (
	*mcp.CallToolResult,
	Output,
	error,
) {
	resources := []ManifestoResource{
		{
			Title:       "The Web Brick Manifesto Overview",
			Description: "A declaration of simplicity and sanity in web development",
			Content: `The Web Brick Manifesto is a philosophy of radical simplicity inspired by Lego blocks. 
We propose treating web components like Lego blocks - simple, self-contained, and predictable.

The core philosophy: The web was meant to be a system of linked documents, and we believe in returning to that core strength. 
Our methods must evolve to make the web simpler and faster, trading needless complexity for clarity, and fragility for robustness.`,
		},
		{
			Title:       "The Three Core Principles",
			Description: "The fundamental rules of the Web Brick methodology",
			Content: `Principle I: One Block, One Identity
Every functional piece of your interface is a unique block represented by a single HTML element with a single, descriptive class name.
Example: <div class="featured-blog-post">...</div>

Principle II: One Identity, One Rule
For every unique class, there shall be one—and only one—corresponding rule in our stylesheet that defines the block's entire appearance.

Principle III: No Mutations. Create Anew.
If an element needs to look even slightly different, it is a different element. It deserves its own identity and its own rule.
Don't modify existing blocks - create new ones instead.`,
		},
		{
			Title:       "The Promise & Benefits",
			Description: "What you gain by following the Web Brick methodology",
			Content: `By adhering to these principles, we gain predictability, portability, and performance:

• Copy & Paste with Confidence: Take an HTML block from one project, copy its single CSS rule, paste them into a new project. It will look identical. Every time.

• Delete without Fear: Need to remove a component? Delete the HTML and its corresponding CSS rule. No side effects.

• Trivial Pruning: Removing unused CSS is no longer complex. If a block isn't used, its single rule can be deleted with absolute certainty.

• Effortless Refactoring: Your codebase becomes a library of well-defined, independent parts.`,
		},
		{
			Title:       "Philosophy & Scope",
			Description: "When and why to use the Web Brick methodology",
			Content: `This methodology is a return to the web's roots of simplicity. It is ideal for:
• Content-rich sites
• Blogs
• Portfolios  
• Marketing pages
• Any project where clarity and performance are paramount

It is intentionally not aimed at solving the unique challenges of highly complex, state-driven web applications. 
We are not trying to build a spaceship with building blocks; we are trying to build elegant, robust, and beautiful houses.

This is not a framework. It is a discipline that can be applied in any environment, with any technology stack.`,
		},
	}

	return nil, Output{
		Message:   "Thank you for using the Web Brick Manifesto tool! Here are resources explaining our methodology for simpler, more maintainable web development.",
		Resources: resources,
	}, nil
}

func main() {
	// Create a server with the web-brick-manifesto tool
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "web-brick-manifesto", 
		Version: "v1.0.0",
	}, nil)
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "web-brick-manifesto",
		Description: "Get information about the Web Brick Manifesto - a methodology for simpler, more maintainable web development inspired by Lego blocks",
	}, WebBrickManifesto)
	
	// Run the server over stdin/stdout, until the client disconnects
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
