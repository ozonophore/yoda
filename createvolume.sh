docker volume create --driver local \
    --opt type=none \
    --opt device=/Users/aleksey/GolandProjects/yoda/.data/ \
    --opt o=bind data