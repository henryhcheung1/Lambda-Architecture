union UserID {
    1: i64 user_id;
}

struct VideoID {
    1: i64 video_id;
}

struct VideoProperty {
    1: required VideoID id;
    2: required VideoPropertyValue property;
}

union VideoPropertyValue {
    1: i32 video_views;
}

struct VideoViewEdge {
    1: required UserID user;
    2: required VideoID video;
    3: required i64 nonce;
}

struct Location {
    1: optional string city;
    2: optional string state;
    3: optional string country;
}

enum GenderType {
    MALE = 1,
    FEMALE = 2
}

union UserPropertyValue {
    1: string full_name;
    2: GenderType gender;
    3: Location location;
}

struct UserProperty {
    1: required UserID id;
    2: required UserPropertyValue property;
}

union DataUnit {
    1: UserProperty user_property;
    2: VideoProperty video_property;
    4: VideoViewEdge video_view;
}

struct TimeStamp {
    1: required i64 true_as_of_secs;
}

struct Data {
    1: required TimeStamp time;
    2: required DataUnit dataunit;
}
