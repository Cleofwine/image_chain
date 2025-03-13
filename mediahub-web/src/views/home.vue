<script lang="ts" setup>
import logo from "./components/logo.vue";
import upload from "./components/upload.vue";
import { ref } from 'vue'
import { onBeforeMount } from 'vue'
import { home } from '../api/api.ts'


let data = {
    banners: ref(["", "", ""]),
    imgs1: ref(["", "", "", "", ""]),
    imgs2: ref(["", "", "", "", ""]),
}

class homeRes {
    banners: Array<string>;
    image1: Array<string>;
    image2: Array<string>;
    constructor(banners: Array<string>,image1: Array<string>,image2: Array<string>) {
        this.banners = banners
        this.image1 = image1
        this.image2 = image2
    }
}

onBeforeMount(() => {
    console.log("startmount")
    home<homeRes>().then(function (res) {
        console.log(res)
        data.banners.value = res.data?.banners || ["", "", ""]
        data.imgs1.value = res.data?.image1 || ["", "", "", "", ""]
        data.imgs2.value = res.data?.image2 || ["", "", "", "", ""]
        console.log(data)
    }).catch(function (res) {
        console.log(res)
    })
})

function getStyle(index: number) {
    let style = "position:absolute;width:14rem;height:10rem;top:0;background-color:#d3dce6;"
    style += " left:" + 15.25 * index + "rem;"
    return style
}

function getItemstyle(url: string) {
    let style = "width:100%;height:100%;"
    style += "background-image:url(" + url + ");"
    style += "background-position:center center;"
    style += "background-repeat:no-repeat;"
    style += "background-size:cover;"
    return style
}
</script>
<template>
    <div style="position: relative; width: 100%; height: 57.1875rem;">
        <el-carousel width=" 100%" height="25.9375rem" class="banner" :interval="5000" arrow="always">
            <el-carousel-item v-for="item in data.banners.value" :key="item">
                <div :style="getItemstyle(item)"></div>
            </el-carousel-item>
        </el-carousel>
        <div class="banner_upper">
            <logo style="position: absolute; top: 0; "></logo>
            <upload style="position: absolute; top: 7.663rem"></upload>
        </div>
        <div
            style="width: 100%; height: 21.875rem; position: relative; top: 5rem; display: flex; justify-content: center; ">
            <div style="position: relative;width: 75rem;height: 21.875rem;">
                <div style="position: relative;width: 100%; height: 10rem">
                    <div v-for="(item, index) in data.imgs1.value" :style="getStyle(index)">
                        <div :style="getItemstyle(item)"></div>
                    </div>
                </div>
                <div style="position: relative; width: 100%; height: 10rem; top: 1.875rem">
                    <div v-for="(item, index) in data.imgs2.value" :style="getStyle(index)">
                        <div :style="getItemstyle(item)">
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.el-carousel__item h3 {
    color: #475669;
    opacity: 0.75;
    line-height: 580px;
    margin: 0;
    text-align: center;
}

.el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
    background-color: #d3dce6;
}

.banner {
    width: 100%;
    min-height: 25.9375rem;
    position: relative;
}

.banner_upper {
    display: flex;
    width: 46.9%;
    height: 11.5rem;
    z-index: 1;
    top: 6rem;
    left: 26.8%;
    position: absolute;
    text-align: center;
    justify-content: center;
}
</style>