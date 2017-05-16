<style>
.xh_loginForm .el-form-item.is-required .el-form-item_label:before{
	content: '';
	margin-right: 0px;
}
.xh_loginForm .title{
	margin:0px auto 30px auto;
	text-align: center;
	color:#888;
}
</style>
<template>
	<el-form :model="xhForm" :rules="rules" ref="xhForm" label-width="100px" class="xh_loginForm">
		<h3 class="title">后台登录</h3>
		<el-form-item label="用户名" prop="username">
			<el-input v-model="xhForm.username"></el-input>
		</el-form-item>
		<el-form-item label="密码" prop="password">
			<el-input type="password" v-model="xhForm.password"></el-input>
		</el-form-item>
		<el-form-item label="验证码" prop="captcha" style="position: relative;">
			<img :src="imgurl" style="position:absolute;top:2px;right:2px;z-index:999;" />
			<el-input v-model="xhForm.captcha" :maxlength="15"></el-input>
			<input type="hidden" v-model="xhForm.captcha_id" />
		</el-form-item>
		<el-form-item>
			<el-button id="loginsubmit" type="primary" @click="submitForm('xhForm')" style="width:100%" v-loading.fucllscreen.lock="fullscreenloading">登录</el-button>
		</el-form-item>
	</el-form>
</template>
<script>
export default{
	props:{
		captchaid:{
			type:String,
			default:'text'
		},
		imgurl:{
			type:String,
			default:'text'
		}
	},
	data(){
		return{
			fullscreenloading:false,
			xhForm:{
				username:'',
				password:'',
				captcha:'',
				captchaid:'',
			},
			rules:{
				username:[
					{require:true,message:'请输入用户名',trigger:'blur'},
					{min:5,max:15,message:'长度在5到15个字符',trigger:'blur'}
				],
				password:[
					{require:true,message:'请输入密码',trigger:'blur'},
					{min:8,max:15,message:'长度在8到15个字符',trigger:'blur'}
				],
			}
		}
	},
	methods:{
		submitForm(formName){
			this.$refs[formName].validate((valid)=>{
				if(valid){
					var e=this;
					$.ajax({
						type:'post',
						url:'/admin',
						data:{username:this.xhForm.username,password:this.xhForm.password,captcha:this.xhForm.captcha,captcha_id:this.captchaid},
						dataType:"json",
						beforeSend:function(){
							e.fullscreenloading =true;
						},
						success:function(msg){
							if(msg.status===4){
								e.fullscreenloading = false;
								e.$message.error(msg.info);
							}else if(msg.status===1){
								e.fullscreenloading = false;
								e.$message.success(msg.info);
								setTimeout(function(){window.location.href="/admin/index"},3000);
							}else{
								e.fullscreenloading = false;
								e.$message.error('网络连接异常！');
							}
						},
						error:function(XMLHttpRequest,testStatus,errorThrown){
							e.fullscreenloading = false;
								e.$message.error('网络连接异常！');
						}
					});
				}else{
					console.log('error submit!');
					return false;
				}
			});
		}
	}
}
</script>