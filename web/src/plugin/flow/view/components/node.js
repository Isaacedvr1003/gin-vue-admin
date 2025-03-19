import { register } from '@antv/x6-vue-shape'

import start from '@/plugin/flow/view/components/nodes/start.vue'
import end from '@/plugin/flow/view/components/nodes/end.vue'
import event from '@/plugin/flow/view/components/nodes/event.vue'
import decision from '@/plugin/flow/view/components/nodes/decision.vue'
// import auto from '@/plugin/flow/view/components/nodes/auto.vue'
import parallel from '@/plugin/flow/view/components/nodes/parallel.vue'
import merge from '@/plugin/flow/view/components/nodes/merge.vue'

const nodeBaseData = {
    title:"",
    decision:"",
    logic:""
}

const ports = {
    groups: {
        top: {
            position: 'top',
            attrs: {
                circle: {
                    r: 4,
                    magnet: true,
                    stroke: '#5F95FF',
                    strokeWidth: 1,
                    fill: '#fff',
                    style: {
                        visibility: 'hidden',
                    },
                },
            },
        },
        right: {
            position: 'right',
            attrs: {
                circle: {
                    r: 4,
                    magnet: true,
                    stroke: '#5F95FF',
                    strokeWidth: 1,
                    fill: '#fff',
                    style: {
                        visibility: 'hidden',
                    },
                },
            },
        },
        bottom: {
            position: 'bottom',
            attrs: {
                circle: {
                    r: 4,
                    magnet: true,
                    stroke: '#5F95FF',
                    strokeWidth: 1,
                    fill: '#fff',
                    style: {
                        visibility: 'hidden',
                    },
                },
            },
        },
        left: {
            position: 'left',
            attrs: {
                circle: {
                    r: 4,
                    magnet: true,
                    stroke: '#5F95FF',
                    strokeWidth: 1,
                    fill: '#fff',
                    style: {
                        visibility: 'hidden',
                    },
                },
            },
        },
    },
    items: [
        {
            group: 'top',
        },
        {
            group: 'right',
        },
        {
            group: 'bottom',
        },
        {
            group: 'left',
        },
    ],
}

register({
    shape: 'start-node',
    width: 48,
    height: 48,
    component: start,
    data:nodeBaseData,
    ports
})

register({
    shape: 'end-node',
    width: 48,
    height: 48,
    component: end,
    data:nodeBaseData,
    ports
})

register({
    shape: 'event-node',
    width: 70,
    height: 48,
    component: event,
    data:nodeBaseData,
    ports
})


// register({
//     shape: 'auto-node',
//     width: 66,
//     height: 36,
//     data:nodeBaseData,
//     component: auto,
//     ports
// })

register({
    shape: 'decision-node',
    width: 56,
    height: 56,
    data:nodeBaseData,
    component: decision,
    ports:ports
})

register({
    shape: 'parallel-node',
    width: 56,
    height: 56,
    data:nodeBaseData,
    component: parallel,
    ports:ports
})


register({
    shape: 'merge-node',
    width: 56,
    height: 56,
    data:nodeBaseData,
    component: merge,
    ports:ports
})


export const RegisterNodes = (graph,stencil) => {
    const decision = graph.createNode({
        shape: 'decision-node',
        data: {
            name: 'decision'
        },
        label: '决策',
    })

    const parallel = graph.createNode({
        shape: 'parallel-node',
        data: {
            name: 'parallel'
        },
        label: '并行',
    })

    const marge = graph.createNode({
        shape: 'merge-node',
        data: {
            name: 'merge'
        },
        label: '合并',
    })

    const fun = graph.createNode({
        shape: 'event-node',
        label: '事件',
        data: {
            name: 'fun',
        },
    })
    const end = graph.createNode({
        shape: 'end-node',
        data: {
            name: 'end',
        },
        label: '结束',
    })


    // const auto = graph.createNode({
    //     shape: 'auto-node',
    //     data:{
    //         name: 'dataSource'
    //     },
    //     attrs: {
    //         body: {
    //             refPoints: '10,0 40,0 30,20 0,20',
    //         },
    //     },
    //     label: '自动',
    // })


    const start = graph.createNode({
        shape: 'start-node',
        label: '开始',
        data:{
            name: 'start'
        },
    })

    stencil.load([start, fun, end], 'node')
    stencil.load([decision,parallel,marge], 'ctrl')

}
