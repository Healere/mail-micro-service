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
                            <el-descriptions-item label="Username:">{{ userInfo.name }}
                                &ensp;&ensp;&ensp;&ensp;</el-descriptions-item>
                            <el-descriptions-item label="Email:">{{ userInfo.email }} </el-descriptions-item>
                        </el-descriptions>

                        <!-- search -->
                        <span style="position: absolute;left: 100vh;top:3vh;width: 50vh;">
                            <el-input v-model="input" class="w-50 m-2" placeholder="Search" :prefix-icon="Search" />
                        </span>
                    </span>
                </div>
                <!-- header-border -->
                <div class="header-layout-line" />
                <br>


            </el-header>

            <el-container>
                <!-- Aside -->
                <el-aside width="20%">
                    <!-- left-layout -->

                    <div>
                        <el-row class="left-layout-row">
                            <el-col class="left-layout-col">
                                <el-button class="left-layout-write-btn" @click="this.$router.push({ path: '/mail' })">
                                    <el-icon>
                                        <Back />
                                    </el-icon>
                                    Back
                                </el-button>
                            </el-col>
                        </el-row>
                    </div>
                </el-aside>

                <!-- Main -->
                <el-main>
                    <div class="writeEmail-box">
                        <el-form ref="ruleFormRef" label-width="10%" class="demo-ruleForm" style="max-width: 90%;"
                            status-icon>
                            <br>
                            <el-form-item label="收件人">
                                <el-input v-model="To" />
                            </el-form-item>
                            <el-form-item label="主题">
                                <el-input v-model="Subject" />
                            </el-form-item>
                            <el-form-item label="正文">
                                <el-input v-model="Text" type="textarea" :rows="20"></el-input>
                            </el-form-item>
                            <br>
                            <el-form-item>
                                <el-button type="primary" @click="send">
                                    Send
                                </el-button>
                                <el-button @click="reset">Reset</el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>
  
<script lang="ts" setup>
import { Search } from '@element-plus/icons-vue'
import { ref } from 'vue'
import { useState } from '../hooks/useState'
import { useActions } from '../hooks/useActions';
const input = ref('')
const To = ref('')
const Subject = ref('')
const Text = ref('')

const { userInfo }: any = useState({
    userInfo: (state: any) => state.userModule.userInfo
})

const storeActions: any = useActions({
    sendEmail: 'emailModule/send',
})
const send = () => {
    const { sendEmail } = storeActions
    const fromAddress = userInfo.value.email
    const fromName = userInfo.value.name
    const to = To.value
    const subject = Subject.value
    const text = Text.value
    sendEmail({ fromAddress, fromName, to, subject, text })
}

const reset = () => {
    Text.value = ""
}

</script>
  
<style scoped>
.writeEmail-box {
    background-color: #ecf5ff;
    width: 90%;
    height: 100%;
}

.writeEmail-box el-input {
    width: 10vh
}

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
  