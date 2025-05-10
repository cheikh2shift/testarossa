import * as vscode from 'vscode';
// @ts-ignore
import { exec } from 'child_process';

export async function activate(context: vscode.ExtensionContext) {

  const config = vscode.workspace.getConfiguration('testarossaContextMenu');
  let geminiKey = config.get<string>('geminiKey') || '';

  exec("go install github.com/cheikh2shift/testarossa/cmd/test-gen@latest", 
    (error : any, stdout : string, stderr : string) => {
        if (error) {
            vscode.window.showErrorMessage(`Error installing test-gen command: ${stderr || error.message}`);
            return;
        }
        vscode.window.showInformationMessage("test-gen installed");
  });
  // Prompt once for GEMINI_KEY if not set
  if (!geminiKey) {
    const input = await vscode.window.showInputBox({
      prompt: 'Enter your GEMINI_KEY (Gemini API key)',
      ignoreFocusOut: true
    });
    if (input) {
        try {
            geminiKey = input;
            await config.update('geminiKey', input, vscode.ConfigurationTarget.Global);
            
            vscode.window.showInformationMessage('GEMINI_KEY saved to settings.');
        } catch (err) {
            vscode.window.showErrorMessage(
                'Unable to save GEMINI_KEY programmatically. Please add it to your settings.json under "goContextMenu.geminiKey".'
            );
        }
    } else {
      vscode.window.showWarningMessage('GEMINI_KEY not set; API features disabled.');
    }
  }
  const disposable = vscode.commands.registerCommand('testarossaContextMenu.customAction', (uri: vscode.Uri) => {
    // If invoked from explorer, uri is provided; otherwise use active editor
    const fileUri = uri || vscode.window.activeTextEditor?.document.uri;
    if (!fileUri) {
      vscode.window.showErrorMessage('No Go file selected.');
      return;
    }
    const filePath = fileUri.fsPath;

    // Compute filename without .go extension
    const fileName = filePath.split('.go')[0];
    vscode.window.showInformationMessage(`Filename (sans .go): ${fileName}`);

    // Get project root
    const workspaceFolder = vscode.workspace.getWorkspaceFolder(fileUri);
    const projectRoot = workspaceFolder ? workspaceFolder.uri.fsPath : vscode.workspace.rootPath;
    vscode.window.showInformationMessage(`Project root path: ${projectRoot}`);

    exec(`GEMINI_KEY=${geminiKey} test-gen -projectroot ${projectRoot} -file ${fileName}.go -output ${fileName}_test.go`, 
        (error : any, stdout : string, stderr : string) => {
            if (error) {
                vscode.window.showErrorMessage(`Error generating tests: ${stderr || error.message}`);
                return;
            }
            vscode.window.showInformationMessage("tests generated");
        }
    );
    
  });

  context.subscriptions.push(disposable);
}

export function deactivate() {}