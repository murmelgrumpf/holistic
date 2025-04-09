let globalSheets = null;

function getStyleSheets() {
    if (globalSheets === null) 
        globalSheets = Array.from(document.styleSheets)
            .map(x => {
                const sheet = new CSSStyleSheet()
                const css = Array.from(x.cssRules).map(rule => rule.cssText).join(' ')
                sheet.replaceSync(css)
                return sheet
            })

    return globalSheets
}

function styleShadowRoot(shadowRoot) {
    shadowRoot.adoptedStyleSheets.push(
        ...getStyleSheets()
    );
    return shadowRoot
}
