
function assert(condition, message, ...args) {
    if (!condition){
        console.log(...args)
        throw Error('Assert failed: ' + (message || ''))
    }
}

function assertNumber(num, name) {
    assert(num !== null && num !== undefined && !isNaN(Number(num)), `${name} is not a number`, `${name}: `, num)
    return Number(num)
}

function assertBool(bool, name) {
    assert(bool !== null && bool !== undefined && (typeof bool === "boolean" || bool === "true" || bool === "false") , `${name} is not a boolean`, `${name}: `, bool)
    return bool === true || bool === "true"
}
