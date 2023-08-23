import {
  Button,
  message,
  ConfigProvider,
  Menu,
  Input,
  Row,
  Col,
  Card,
  Icon,
  Space,
  Avatar,
  List,
  Tag,
} from 'ant-design-vue'

import Vue from 'vue'

// 引入自定义icon
import iconFront from '../assets/aliicons/iconfont'
const aliicon = Icon.createFromIconfontCN({
  scriptUrl: iconFront
})
Vue.component('ali-icon', aliicon)

message.config({
  top: `60px`,
  duration: 2,
  maxCount: 3,
});
Vue.prototype.$message = message

Vue.use(ConfigProvider)
Vue.use(Menu)
Vue.use(Button)
Vue.use(Input)
Vue.use(Row)
Vue.use(Col)
Vue.use(Card)
Vue.use(Icon)
Vue.use(Space)
Vue.use(Avatar)
Vue.use(List)
Vue.use(Tag)