export const RegisterEvent = (graph,graphContainerRef,event) => {
    graph.bindKey(['meta+c', 'ctrl+c'], () => {
        event['copy']()
        return false
    })
    graph.bindKey(['meta+x', 'ctrl+x'], () => {
        event['cut']()
        return false
    })
    graph.bindKey(['meta+v', 'ctrl+v'], () => {
        event['paste']()
        return false
    })

// undo redo
    graph.bindKey(['meta+z', 'ctrl+z'], () => {
        event['undo']()
        return false
    })
    graph.bindKey(['meta+shift+z', 'ctrl+shift+z'], () => {
        event['redo']()
        return false
    })

// select all
    graph.bindKey(['meta+a', 'ctrl+a'], () => {
        event['selectAll']()
    })

// delete
    graph.bindKey(['backspace','Delete'], () => {
        event['delete']()
    })

// zoom
    graph.bindKey(['ctrl+1', 'meta+1'], () => {
        const zoom = graph.zoom()
        if (zoom < 1.5) {
            graph.zoom(0.1)
        }
    })
    graph.bindKey(['ctrl+2', 'meta+2'], () => {
        const zoom = graph.zoom()
        if (zoom > 0.5) {
            graph.zoom(-0.1)
        }
    })

// 控制连接桩显示/隐藏
    const showPorts = (ports, show) => {
        for (let i = 0, len = ports.length; i < len; i += 1) {
            ports[i].style.visibility = show ? 'visible' : 'hidden'
        }
    }
    graph.on('node:mouseenter', () => {
        const container = graphContainerRef.value
        const ports = container.querySelectorAll(
            '.x6-port-body',
        )
        showPorts(ports, true)
    })
    graph.on('node:mouseleave', () => {
        const container = graphContainerRef.value
        const ports = container.querySelectorAll(
            '.x6-port-body',
        )
        showPorts(ports, false)
    })

    graph.on('cell:dblclick', (e) => {
        event['cell:dblclick'](e)
    })
}
