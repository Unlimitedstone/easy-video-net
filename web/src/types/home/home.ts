import { PageInfo } from "../idnex"

//getLiveRoom 获取直播接口请求需要
export interface GetLiveRoomRes {
    address: string,
    key: string
}

//轮播图
export interface Rotograph {
    title: string
    cover: string
    color: string
    type: string
    to_id: number
}
export type RotographList = Array<Rotograph>

//视频信息
export interface VideoInfo {
    id: number
    uid: number
    title: string
    video: string
    cover: string
    video_duration: number
    label: Array<string>
    introduce: string
    heat: number
    barrageNumber: number
    username: string
    created_at: string
}

export interface MovieMessages {
    id: number;
    uni_id: string;
    name: string;
    image_url: string;
    image_src: string
    score: string;
    type: string;
    common: string;
    source: string;
    create_time: string;
    update_time: string;
    introduce: string;
    alias: string;
    duration: number;
    release_date: string;
    director: string;
    actor: string;
    writer: string;
    country: string;
    imdb: string;
}

export type VideoInfoList = Array<VideoInfo>

export interface GetHomeInfoReq {
    page_info: PageInfo
}

export interface GetHomeInfoRes {
    rotograph: RotographList
    videoList: VideoInfoList
    movie_messages: Array<MovieMessages>
}