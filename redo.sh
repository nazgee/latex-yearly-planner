#!/bin/bash
rm ./out/* -f && PLANNER_YEAR=2023 PASSES=1 CFG="cfg/base.yaml,cfg/rm2.base.yaml,cfg/template_breadcrumb.yaml,cfg/rm2.breadcrumb.default.yaml,cfg/rm2.breadcrumb.default.dailycal.yaml" NAME=kalendarz ./single.sh && evince ./kalendarz.pdf
