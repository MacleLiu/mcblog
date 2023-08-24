<template>
    <a-row class="headerRow">

        <a-col :xs="4" :sm="6" :md="6" :lg="4">
            <div class="logoContainer">
                <img style="width: 40px;" src="../../assets/mc.svg" @click="toIndex"/>
            </div>
        </a-col>

        <a-col :xs="14" :sm="12" :md="14" :lg="8">
            <div class="inputContainer">
                <a-input-search placeholder="搜索" style="max-width: 200px; width: auto; min-width: 0; flex: auto;" @search="onSearch" />
            </div>
        </a-col>

        <a-col :xs="6" :sm="6" :md="4" :lg="12" >
            <div class="menuContainer">
                <a-menu
                    mode="horizontal"
                    class="antMenu"
                    @click="goToPage"
                    :overflowedIndicator = "menuIcon"
                >
                    <a-menu-item key="pigeonhole">
                        归档
                    </a-menu-item>
                    <a-menu-item key="work">
                        作品
                    </a-menu-item>
                    <a-menu-item key="tool">
                        工具
                    </a-menu-item>
                    <a-menu-item key="about">
                        关于
                    </a-menu-item>
                </a-menu>
            </div>
        </a-col>
    </a-row>
</template>

<script>

    export default {
        methods: {
            menuIcon() {
                return <a-icon type="menu" />
            },
            toIndex() {
                this.$router.push('/index').catch((err) => err )
            },
            goToPage(item) {
                this.$router.push('/' + item.key).catch((err) => err )
            },
            onSearch() {},
        }
    }
</script>

<style scoped>
    .headerRow {
        height: 6%;
        flex: 0 0 auto;
        display: flex;
        padding: 0 40px;
        box-shadow: 0px 8px 8px -8px rgb(200, 200, 200);
    }
    .logoContainer {
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: left;
        font-size: 32px;
        color: black;
        font-weight: bold;
    }
    .logoContainer img {
        margin-right: 10px;
    }
    .inputContainer {
        height: 100%;
        width: 100%;
        display: flex;
        align-items:center;
        justify-content: left;
    }
    .menuContainer {
        height: 100%;
        width: 100%;
        display: flex;
        align-items: flex-end;
        justify-content: flex-end;
    }
    .antMenu {
        text-align: right;
        font-size: 20px;
        font-family: serif;
        font-weight: bold;
        color: rgb(80, 80, 80);
        min-width: 0;
        flex: auto;
    }
</style>