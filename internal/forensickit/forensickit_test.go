// internal/forensickit/forensickit_test.go
package forensickit

import (
    "testing"
)

// TestNewApp tests the NewApp function with verbose logging enabled.
func TestNewApp(t *testing.T) {
    app := NewApp(true)
    if app == nil {
        t.Fatal("NewApp returned nil")
    }
    if !app.Verbose {
        t.Errorf("Expected verbose to be true, but got %v", app.Verbose)
    }
    if app.ProcessedCount != 0 {
        t.Errorf("Expected ProcessedCount to be 0, got %d", app.ProcessedCount)
    }
}

// TestProcess tests the Process function with a sample input.
func TestProcess(t *testing.T) {
    app := NewApp(false)
    result, err := app.Process("test data")
    
    if err != nil {
        t.Errorf("Process returned error: %v", err)
    }
    
    if !result.Success {
        t.Errorf("Expected result.Success to be true, but got %v", result.Success)
    }
    
    if app.ProcessedCount != 1 {
        t.Errorf("Expected ProcessedCount to be 1, got %d", app.ProcessedCount)
    }
}

// TestRun tests the Run function with empty arguments.
func TestRun(t *testing.T) {
    app := NewApp(false)
    err := app.Run("", "")
    
    if err != nil {
        t.Errorf("Run returned error: %v", err)
    }
}