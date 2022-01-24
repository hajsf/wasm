async function init(){
    const WASM_URL = 'main.wasm';
    const go = new Go();
    const result = await WebAssembly.instantiateStreaming(
        fetch(`http://localhost:8080/www/wasm/app/${WASM_URL}`),
        go.importObject
       /* go.importObject.env = {
            'main.add': function(x, y) {
                return x + y
            }
            // ... other functions
        } */
    );
    go.run(result.instance); 
}