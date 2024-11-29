{
    if($1 == "www.jrock.us/") {
        next
    }
    sub(/www.jrock.us/, "srv")
    sub(/uid=[0-9\.]+/, "uid=65534")
    sub(/gid=[0-9\.]+/, "gid=65534")
    sub(/mode=0755 type=file/, "mode=0644 type=file")
    print;
}
