<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="汽修店名称:">
                <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="位置:">
                <el-input v-model="formData.location" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="营业时间 :">
                <el-input v-model="formData.openingHours" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="经营范围:">
                <el-input v-model="formData.businessScope" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="联系人:">
                <el-input v-model="formData.contacter" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="联系电话:">
                <el-input v-model="formData.contact" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="店铺详情:">
                <el-input v-model="formData.detail" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createGarage,
    updateGarage,
    findGarage
} from "@/api/xiu_garage";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Garage",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            name:"",
            location:"",
            openingHours:"",
            businessScope:"",
            contacter:"",
            contact:"",
            detail:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createGarage(this.formData);
          break;
        case "update":
          res = await updateGarage(this.formData);
          break;
        default:
          res = await createGarage(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findGarage({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.regarage
       this.type == "update"
     }
    }else{
     this.type == "create"
   }
  
}
};
</script>

<style>
</style>