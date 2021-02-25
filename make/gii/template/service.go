/**
 * @Time : 2020/10/9 17:32
 * @Author : liang
 * @File : service
 * @Software: GoLand
 */

package template

const Service = `<?php


namespace %s;


interface %sService 
{

}
`

const ServiceImpl = `<?php


namespace %s\impl;

use %sService;

class %sServiceImpl implements %sService 
{

}
`
