import { writable } from 'svelte/store'

function createUniqueArray() {
    const { subscribe, update } = writable([])

    return {
        subscribe,

        add(item) {
            update(arr => {
                if (arr.indexOf(item) !== -1) {
                    return arr
                }
                const ret = [...arr, item]
                // ret.sort()
                return ret
            })
        }
    }
}

function createString(def = '') {
    const { subscribe, set, update } = writable(def)

    return {
        subscribe,

        clear(v = '') {
            set(v)
        },
        add(msg = '') {
            update(str => str + msg + '\n')
        }
    }
}

export const knownFiles = createUniqueArray()
export const fileContents = createString('Nenhum arquivo selecionado.')
