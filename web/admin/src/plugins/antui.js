import { 
  Button,
  FormModel,
  Input,
  Icon,
  message,
  Layout,
  Popconfirm,
  Menu,
  Card,
  Table,
  Row,
  Col,
  ConfigProvider,
  Modal,
  Select,
  Upload,
  Switch,
  Tag,
} from 'ant-design-vue'

import Vue from 'vue'

message.config({
    top: `60px`,
    duration: 2,
    maxCount: 3,
  });
Vue.prototype.$message = message
Vue.prototype.$confirm = Modal.confirm

Vue.use(Button)
Vue.use(FormModel)
Vue.use(Input)
Vue.use(Icon)
Vue.use(Layout)
Vue.use(Popconfirm)
Vue.use(Menu)
Vue.use(Card)
Vue.use(Table)
Vue.use(Row)
Vue.use(Col)
Vue.use(ConfigProvider)
Vue.use(Modal)
Vue.use(Select)
Vue.use(Upload)
Vue.use(Switch)
Vue.use(Tag)