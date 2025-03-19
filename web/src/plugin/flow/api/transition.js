import service from "@/utils/request";


export const startTransitionExample = (data) => {
    return service({
        url: '/flow/startTransitionExample',
        method: 'post',
        data
    })
}


export const getTransitionExampleStep = (params) => {
    return service({
        url: '/flow/getTransitionExampleStep',
        method: 'get',
        params
    })
}


export const getTransitionExampleList = (params) => {
    return service({
        url: '/flow/getTransitionExampleList',
        method: 'get',
        params
    })
}

export const getTransitionExampleNext = (data) => {
    return service({
        url: '/flow/transitionExampleNext',
        method: 'post',
        data
    })
}
