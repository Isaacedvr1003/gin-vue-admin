<template>
  <div class="h-full flex flex-col">
    <div class="flex justify-between mb-2">
      <div class="gva-flow-tool-bar">
        <span @click="copy">复制(ctrl+c)</span>
        <span @click="cut">剪切(ctrl+x)</span>
        <span @click="paste">粘贴(ctrl+v)</span>
        <span @click="undo">撤销(ctrl+z)</span>
        <span @click="redo">重做(ctrl+shift+z)</span>
        <span @click="del">删除(delete)</span>
        <span @click="selectAll">全选(ctrl+a)</span>
        <span @click="exportSVG">导出SVG</span>
        <span @click="exportPNG">导出PNG</span>
        <span @click="exportJPG">导出JPG</span>
      </div>
      <el-button type="primary" @click="save">
        保存
      </el-button>
    </div>
    <div class="h-full flex gap-2">
      <div ref="stencilRef" class="w-52 h-full relative bg-white border border-solid border-gray-200" />
      <div ref="graphContainerRef" class="flex-1 h-full bg-white border border-solid border-gray-200" />
      <el-drawer
          title="节点属性"
          v-model="drawerVisible"
          direction="rtl"
          size="40%"
          :close-on-press-escape="false"
          :show-close="false"
      >
        <el-form label-position="top">
          <el-form-item label="节点ID">
            {{nodeData.nodeId}}
          </el-form-item>
          <el-form-item label="节点类型">
            {{nodeData.name}}
          </el-form-item>
          <el-form-item label="节点名称">
            <el-input v-model="nodeData.title" />
          </el-form-item>
          <el-form-item label="节点描述">
            <el-input v-model="nodeData.description" />
          </el-form-item>
          <el-form-item label="逻辑值">
            <el-input v-model="nodeData.logic" />
          </el-form-item>
        </el-form>

        <template #header>
          <div class="flex justify-between items-center">
            <span class="text-lg">节点属性</span>
            <div>
              <el-button @click="closeDialog">取 消</el-button>
              <el-button
                  type="primary"
                  @click="enterDialog"
              >确 定</el-button>
            </div>
          </div>
        </template>


      </el-drawer>
    </div>
  </div>

</template>

<script setup>

import {RegisterNodes} from "@/plugin/flow/view/components/node";

defineOptions({
  name: 'FlowGraph'
})


const props = defineProps({
  flowUUID: {
    type: String,
    required: true
  }
})

import {onMounted, ref} from "vue";
import {Graph, Shape} from "@antv/x6";
import {Selection} from "@antv/x6-plugin-selection";
import {Snapline} from "@antv/x6-plugin-snapline";
import {Keyboard} from "@antv/x6-plugin-keyboard";
import {Clipboard} from "@antv/x6-plugin-clipboard";
import {History} from "@antv/x6-plugin-history";
import {Stencil} from "@antv/x6-plugin-stencil";
import { Export } from '@antv/x6-plugin-export'
import {RegisterEvent} from "@/plugin/flow/view/components/event";
import {findGraph, saveGraph} from "@/plugin/flow/api/flow";
import {ElMessage} from "element-plus";

const graphContainerRef = ref()

const graphRef = ref()

const stencilRef = ref()

const drawerVisible = ref(false)


const save = async () => {
  const graphJSON = graphRef.value.toJSON();  // 获取node属性
  const req = {}
  req.flowUUID = props.flowUUID
  const nodes = []
  const edges = []
  graphJSON.cells.forEach((item) => {
      if (item.shape === 'edge') {
        edges.push(item)
      }else{
        nodes.push(item)
      }
  })
  req.nodes = nodes
  req.edges = edges
  const res = await saveGraph(req)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '保存成功'
    })
  }
}

const getGraph = async () => {
  const res = await findGraph({uuid: props.flowUUID})
    if (res.code === 0 && res.data.reflow.ID) {
      const jsonData = {
        cells:[]
      }
      const reflow = res.data.reflow
      reflow.edges.forEach((item) => {
        jsonData.cells.push(item)
      })
      reflow.nodes.forEach((item) => {
        jsonData.cells.push(item)
      })
      console.log(jsonData)
      graphRef.value.fromJSON(jsonData)
    }
}


const copy= () => {
  const cells = graphRef.value.getSelectedCells()
  if (cells.length) {
    graphRef.value.copy(cells)
  }
}
const cut = () => {
  const cells = graphRef.value.getSelectedCells()
  if (cells.length) {
    graphRef.value.cut(cells)
  }
}
const paste = () => {
  if (!graphRef.value.isClipboardEmpty()) {
    const cells = graphRef.value.paste({ offset: 32 })
    graphRef.value.cleanSelection()
    graphRef.value.select(cells)
  }
}

const selectAll = () => {
  const nodes = graphRef.value.getNodes()
  if (nodes) {
    graphRef.value.select(nodes)
  }
}
const undo = () => {
  if (graphRef.value.canUndo()) {
    graphRef.value.undo()
  }
}
const redo = () => {
  if (graphRef.value.canRedo()) {
    graphRef.value.redo()
  }
}
const del = () =>{
  const cells = graphRef.value.getSelectedCells()
  if (cells.length) {
    graphRef.value.removeCells(cells)
  }
}

const nodeData = ref({
  name: '',
  description: '',
  logic:''
})

const nodeClick = async (e) =>{

  const node = e.cell
  const data = node.getData();  // 获取节点的数据
  data.nodeId = node.id
  data.edgeId = node.id
  nodeData.value = data
  drawerVisible.value = true
  // 清除所有的文本选择
  window.getSelection().removeAllRanges();
}

const closeDialog = () => {
  drawerVisible.value = false
}

const enterDialog = () => {
  const node = graphRef.value.getCellById(nodeData.value.nodeId)
  node.setData({
    title: nodeData.value.title,
    description: nodeData.value.description
  })
  drawerVisible.value = false
}

const exportSVG = () => {
  graphRef.value.exportSVG()
}

const exportPNG = () => {
  graphRef.value.exportPNG()
}

const exportJPG = () => {
  graphRef.value.exportJPEG()
}


onMounted(()=>{
  const graph = new Graph({
    container: graphContainerRef.value,
    grid: true,
    mousewheel: {
      enabled: true,
      zoomAtMousePosition: true,
      modifiers: 'ctrl',
      minScale: 0.5,
      maxScale: 3,
    },
    connecting: {
      router: 'manhattan',
      connector: {
        name: 'rounded',
        args: {
          radius: 8,
        },
      },
      anchor: 'center',
      connectionPoint: 'anchor',
      allowBlank: false,
      snap: {
        radius: 20,
      },
      createEdge() {
        return new Shape.Edge({
          attrs: {
            line: {
              stroke: '#1890ff',
              strokeDasharray: 5,
              targetMarker: 'classic',
              style: {
                animation: 'ant-line 45s infinite linear',
              },
            },
          },
          zIndex: 0,
          data:{
            name:"edge",
            edgeId:"",
            title:"",
            descriptor:""
          }
        })
      },
      validateConnection({ targetMagnet }) {
        return !!targetMagnet
      },
    },
  })

  graphRef.value = graph

  const stencil = new Stencil({
    title: '智能体',
    target: graph,
    stencilGraphWidth: 180,
    stencilGraphHeight: 200,
    groups: [
      {
        title: '节点',
        name: 'node',
      },
      {
        title: "逻辑控制",
        name: 'ctrl',
      }
    ],
    layoutOptions: {
      columns: 2,
      columnWidth: 80,
      rowHeight: 65,
    },
  })

  // #region 使用插件
  graph
      .use(
          new Selection({
            rubberband: true,
            showNodeSelectionBox: true,
            showEdgeSelectionBox: true,
          }),
      )
      .use(new Snapline())
      .use(new Keyboard())
      .use(new Clipboard())
      .use(new History())
      .use(new Export())
// #endregion
  stencilRef.value.appendChild(stencil.container)



  // #region 快捷键与事件
  RegisterEvent(graph,graphContainerRef,{
    'cell:dblclick':nodeClick,
    'copy':copy,
    'cut':cut,
    'paste':paste,
    'selectAll':selectAll,
    'undo':undo,
    'redo':redo,
    'delete':del
  })
// #endregion

  // #region 初始化图形
  RegisterNodes(graph, stencil)

  // graph.fromJSON(obj); 给与默认属性

  getGraph()

})

</script>

<style lang="scss">
.x6-widget-stencil-content{
  @apply bg-white;
}
.gva-flow-tool-bar{
  @apply flex gap-1;
  span{
    @apply border border-solid px-2 py-1 rounded cursor-pointer text-sm flex items-center leading-3;
  }
}

@keyframes ant-line {
  to {
    stroke-dashoffset: -1000
  }
}
</style>
