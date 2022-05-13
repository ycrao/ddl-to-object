package com.example.sample.domain.entity;

import com.fasterxml.jackson.databind.PropertyNamingStrategies;
import com.fasterxml.jackson.databind.annotation.JsonNaming;
import lombok.Data;
import java.io.Serializable;
import java.util.Date;

@JsonNaming(PropertyNamingStrategies.SnakeCaseStrategy.class)
@Data
@SuppressWarnings("unused")
public class Article implements Serializable {

    private static final long serialVersionUID = 8648389512580360946L;

    private Long id;

    private Long userId;

    private String content;

    private Date createTime;

    private Date updateTime;
}