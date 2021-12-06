/**
 * @Time : 2020/10/10 11:46
 * @Author : liang
 * @File : reset
 * @Software: GoLand
 */

package template

const Reset = `<?php


namespace %s;


use common\librarys\ApiResponse;
%s

/**
 * @module
 */
class %sController extends %s
{

    /**
     * @return ApiResponse
     */
    public function actionIndex()
    {
        return ApiResponse::success();
    }

    /**
     * @return ApiResponse
     */
    public function actionCreate()
    {
        return ApiResponse::success();
    }

    /**
     * @param $id
     * @return ApiResponse
     */
    public function actionUpdate($id)
    {
        return ApiResponse::success();
    }

    /**
     * @param $id
     * @return ApiResponse
     */
    public function actionDelete($id)
    {
        return ApiResponse::success();
    }
}
`
