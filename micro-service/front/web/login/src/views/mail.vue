<template>
  <!-- contaniner -->
  <div class="common-layout">
    <el-container>
      <el-header height="20vh">
        <!-- header -->
        <div>
          <!-- Icon -->
          <el-avatar class="mr-3" :size="64" src="../src/assets/images/jogIcon.jpg" />
          <!-- UserInfp -->
          <span style="position: absolute;left: 10%;">
            <el-avatar class="mr-3" :size="32" src="../src/assets/images/userIcon.jpg" />
            <span class="text-large font-600 mr-3"> Jog </span>
            <span class="text-sm mr-2" style="color: var(--el-text-color-regular)">
              Email
            </span>
            <el-descriptions :column="3" size="large" class="mt-4">
              <el-descriptions-item label="Username:"> {{ userInfo.name }} &ensp;&ensp;&ensp;&ensp;</el-descriptions-item>
              <el-descriptions-item label="Email:"> {{ userInfo.email }} </el-descriptions-item>
            </el-descriptions>

            <!-- search -->
            <span style="position: absolute;left: 100vh;top:3vh;width: 50vh;">
              <el-input v-model="input" class="w-50 m-2" placeholder="Search" :prefix-icon="Search" />
            </span>
          </span>

          <!-- logout -->
          <span>
            <el-button style="margin-left: 90%;margin-top: -3%;" @click="logout()">
              log out
            </el-button>
          </span>
        </div>
        <!-- header-border -->
        <div class="header-layout-line" />
        <br>

        <span>
          <el-button :icon="Refresh" style="margin-left:20% ;margin-top: -2%;" @click="refresh" />
          <!-- Main-border -->
          <div style="border:1px solid #CDD0D6;width: 100%;margin-left: 20%;margin-top: -0.5%;" />
          <br>
        </span>
      </el-header>

      <el-container>
        <!-- Aside -->
        <el-aside width="20%">
          <!-- left-layout -->

          <div>
            <el-row class="left-layout-row">
              <el-col class="left-layout-col">
                <el-button class="left-layout-write-btn" @click="this.$router.push({ path: '/writeEmail' })">
                  <el-icon>
                    <EditPen />
                  </el-icon>
                  Write
                </el-button>
              </el-col>
            </el-row>
          </div>
        </el-aside>

        <!-- Main -->
        <el-main>

        </el-main>
      </el-container>

    </el-container>


  </div>
</template>

<script lang="ts" setup>
import { Refresh, Search } from '@element-plus/icons-vue'
import { inject, ref } from 'vue'
import {useState} from '../hooks/useState'
import {useActions} from '../hooks/useActions'
import router from '../router/index'
const input = ref('')

const refresh = () => {
  inject("reload");
}

const {userInfo}:any = useState({
  userInfo:(state:any) => state.userModule.userInfo
})

const storeActions:any = useActions({
  logout:'userModule/logout'
})
const logout = () =>{
  const {logout} = storeActions
  logout({})
  router.replace({name:'register'})
}

</script>

<script lang="ts">


export default {
  mounted() {
    if (location.href.indexOf("#reloaded") == -1) {
      location.href = location.href + "#reloaded";
      location.reload();
    }
    
  },
  
};
</script>

<style scoped>
.left-layout-write-btn:hover {
  font-size: 20px;
  border: 0;
  width: 20vh;
  background-color: #CDD0D6;

}

.left-layout-write-btn {
  font-size: 20px;
  border: 0;
  width: 20vh;
  background-color: white;

}

.header-layout-line {
  border: 1px solid #CDD0D6;
  margin-top: 1%;
}

.left-layout-row {
  margin-bottom: 20px;
  width: 10%;
  height: 50px;
}

.left-layout-col {
  border-radius: 100px;
}
</style>
