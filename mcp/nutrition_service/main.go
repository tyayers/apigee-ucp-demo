package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"google.golang.org/genai"
)

func main() {
	transport := flag.String("transport", "streamable", "Transport mode: 'streamable', 'sse', or 'stdio'")
	port := flag.Int("port", 8080, "Port for HTTP/SSE transport")
	flag.Parse()

	// Initialize Gemini Client
	ctx := context.Background()

	// Default values or env vars for Vertex AI
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		log.Fatal("GOOGLE_CLOUD_PROJECT environment variable must be set")
	}
	location := os.Getenv("GOOGLE_CLOUD_LOCATION")
	if location == "" {
		location = "us-central1"
	}

	model := os.Getenv("MODEL")

	clientConfig := &genai.ClientConfig{
		Backend:  genai.BackendVertexAI,
		Project:  projectID,
		Location: location,
	}

	client, err := genai.NewClient(ctx, clientConfig)
	if err != nil {
		log.Fatalf("Failed to create genai client: %v", err)
	}

	// Create MCP server
	s := server.NewMCPServer("nutrition-service", "1.0.0")

	// 1. ingredients-tool
	// s.AddTool(mcp.NewTool("ingredients-tool",
	// 	mcp.WithDescription("Deliver an ingredients list for a given food description."),
	// 	mcp.WithString("food", mcp.Required(), mcp.Description("Description of the food")),
	// ), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 	food := request.Params.Arguments.(map[string]any)["food"].(string)
	// 	prompt := fmt.Sprintf("Give a simple ingredients list for this food, just the list in bullet points: %s", food)
	// 	resp, err := client.Models.GenerateContent(ctx, model, genai.Text(prompt), nil)
	// 	if err != nil {
	// 		return mcp.NewToolResultError(fmt.Sprintf("API error: %v", err)), nil
	// 	}
	// 	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
	// 		return mcp.NewToolResultError("No response from model"), nil
	// 	}
	// 	text := resp.Candidates[0].Content.Parts[0].Text
	// 	return mcp.NewToolResultText(text), nil
	// })

	// 2. gluten-check-tool
	s.AddTool(mcp.NewTool("gluten-check-tool",
		mcp.WithDescription("Return if the food contains gluten (true or false)."),
		mcp.WithString("food", mcp.Required(), mcp.Description("Description of the food")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		food := request.Params.Arguments.(map[string]any)["food"].(string)
		prompt := fmt.Sprintf("Does this food contain gluten? Answer strictly 'true' or 'false': %s", food)
		resp, err := client.Models.GenerateContent(ctx, model, genai.Text(prompt), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %v", err)), nil
		}
		if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
			return mcp.NewToolResultError("No response from model"), nil
		}
		text := resp.Candidates[0].Content.Parts[0].Text
		return mcp.NewToolResultText(text), nil
	})

	// 3. lactose-check-tool
	s.AddTool(mcp.NewTool("lactose-check-tool",
		mcp.WithDescription("Return if the food contains lactose (true or false)."),
		mcp.WithString("food", mcp.Required(), mcp.Description("Description of the food")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		food := request.Params.Arguments.(map[string]any)["food"].(string)
		prompt := fmt.Sprintf("Does this food contain lactose? Answer strictly 'true' or 'false': %s", food)
		resp, err := client.Models.GenerateContent(ctx, model, genai.Text(prompt), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %v", err)), nil
		}
		if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
			return mcp.NewToolResultError("No response from model"), nil
		}
		text := resp.Candidates[0].Content.Parts[0].Text
		return mcp.NewToolResultText(text), nil
	})

	// 4. diabetic-check-tool
	s.AddTool(mcp.NewTool("diabetic-check-tool",
		mcp.WithDescription("Check if the food is ok for people who are diabetic."),
		mcp.WithString("food", mcp.Required(), mcp.Description("Description of the food")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		food := request.Params.Arguments.(map[string]any)["food"].(string)
		prompt := fmt.Sprintf("Is this food OK for people who are diabetic? Provide a brief explanation: %s", food)
		resp, err := client.Models.GenerateContent(ctx, model, genai.Text(prompt), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %v", err)), nil
		}
		if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
			return mcp.NewToolResultError("No response from model"), nil
		}
		text := resp.Candidates[0].Content.Parts[0].Text
		return mcp.NewToolResultText(text), nil
	})

	// Start the server based on the selected transport
	switch *transport {
	case "streamable":
		log.Printf("Starting Streamable HTTP server on port %d...", *port)
		streamableServer := server.NewStreamableHTTPServer(s)
		if err := streamableServer.Start(fmt.Sprintf(":%d", *port)); err != nil {
			log.Fatalf("Server error: %v\n", err)
		}
	case "sse":
		log.Printf("Starting SSE server on port %d...", *port)
		sseServer := server.NewSSEServer(s)
		if err := sseServer.Start(fmt.Sprintf(":%d", *port)); err != nil {
			log.Fatalf("Server error: %v\n", err)
		}
	case "stdio":
		fallthrough
	default:
		if err := server.ServeStdio(s); err != nil {
			fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
			os.Exit(1)
		}
	}
}
